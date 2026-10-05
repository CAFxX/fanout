module mul64(input [63:0] val, input [63:0] key, output [63:0] out);
  assign out = val * key;
endmodule
