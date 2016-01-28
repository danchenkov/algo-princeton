class UF
	def initialize(total = 0)
		@id = []
		(0...total).each {|i| @id[i] = i}
	end

	def connected(p, q)
		@id[p] == @id[q]
	end

	def union(p, q)
		if @id[p] != @id[q]
			pid = @id[p]
			qid = @id[q]
			0.upto(@id.length) do |i|
				@id[i] = qid if @id[i] == pid
			end
		end
	end

	def to_s
		return @id.to_s
	end
end

n = STDIN.readline.to_i
uf = UF.new(n)

while !STDIN.eof
	p, q = STDIN.readline.split(" ").map{|x| x.to_i}
	uf.union(p, q)
end

puts uf
puts "0, 1 connected: #{uf.connected(0, 1)}"
puts "2, 3 connected: #{uf.connected(2, 3)}"

