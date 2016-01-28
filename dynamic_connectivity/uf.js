var Class = function(methods) {
	var klass = function() {
		this.initialize.apply(this, arguments);
	};

	for (var property in methods) {
		klass.prototype[property] = methods[property];
	}

	if (!klass.prototype.initialize) klass.prototype.initialize = function(){};

	return klass;
};

var UF = Class({
	initialize: function(total) {
		this.id = [];
		for (i=0; i<total; i++) { this.id[i] = i }
	}, toString: function() {
		s = "[";
		for (i=0; i<this.id.length; i++) {
			if (i != this.id.length - 1) {
				s = s + this.id[i] + ", ";
			} else {
				s = s + this.id[i];
			}
		}
		s = s + "]";
		return s;
	}, union: function(p, q) {
		pid = this.id[p]
		qid = this.id[q]
		if (pid == qid) return;
		for (i=0; i<this.id.length; i++) {
			if (this.id[i] == pid) {
				this.id[i] = qid
			}
		}
		return;
	}, connected: function(p, q) {
		return this.id[p] == this.id[q]
	}
});

var readline = require('readline');
var rl = readline.createInterface({
	input: process.stdin,
	output: process.stdout,
	terminal: false
});

var uf;
rl.resume();

rl.on('line', function(line){
	if (uf == undefined) {
		uf = new UF(parseInt(line));
	} else {
		p = line.split(" ")[0];
		q = line.split(" ")[1];
		uf.union(p, q);
	}
});

rl.on('close', function(){
	console.log(uf.toString());
	console.log("0, 1 connected:", uf.connected(0, 1));
	console.log("2, 3 connected:", uf.connected(2, 3));
});


