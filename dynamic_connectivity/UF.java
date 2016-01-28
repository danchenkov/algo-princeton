import java.util.Arrays;

public class UF {
	public static void main(String[] args) {
		int N = StdIn.readInt();
		UF uf = new UF(N);

		while (!StdIn.isEmpty()) {
			int p = StdIn.readInt();
			int q = StdIn.readInt();
			if (!uf.connected(p, q)) {
				uf.union(p, q);
			}
		}

		StdOut.println(uf.toString());
		StdOut.println("0, 1 connected: " + uf.connected(0, 1));
		StdOut.println("2, 3 connected: " + uf.connected(2, 3));
	}

	private int[] id;

	public UF(int N) {
		id = new int[N];
		for (int i=0; i<N; i++) {
			id[i] = i;
		}
	}

	public boolean connected(int p, int q) {
		return id[p] == id[q];
	}

	public void union(int p, int q) {
		int pid = id[p];
		int qid = id[q];

		if (pid != qid) {
			for (int i=0; i<id.length; i++) {
				if (id[i] == pid) id[i] = qid;
			}
		}
	}

	public String toString() {
		return Arrays.toString(id);
	}
}
