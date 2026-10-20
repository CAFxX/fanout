package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }

func main() {
    var v1 string = "a"
    v2 := []string{v1, ((v1 + (("" + v1) + v1)) + v1), ((v1 + v1) + "foo bar"), v1}
    v3 := []uint32{65535, 0, 0, 2147483647}
    v4 := make(map[uint32]bool)
    var v5 string = "hello"
    v5 = "hello"
    v5 = (("\n\t" + v1) + "")
    var v6 uint8 = 2
    fmt.Println(v1)
    if (v6 != (v6 << ((322) & 63))) {
        v1 = ("a" + v1)
        for fi7 := 0; fi7 < 16; fi7++ {
            if (!true) {
                _ = len(v2)
            } else {
                switch (v5 + ("a" + "foo bar")) {
                case "foo bar":
                    for fi8 := 0; fi8 < 17; fi8++ {
                        fmt.Println(fi7, v5, fi8)
                        v4[0] = true
                        _ = fi8
                    }
                case "":
                    v9 := GenS1{F0: 65535, F1: 125}
                    _ = v9
                case "hello":
                    if (true || true) {
                        switch 0 {
                        case 32768:
                            v1 = "\n\t"
                        }
                    } else {
                        v6 += (^(v6))
                        v6 = (v6 * v6)
                        for fi10 := 0; fi10 < 18; fi10++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi10
                        }
                    }
                }
                var v11 int = fi7
                _ = v11
            }
            v1 = "\n\t"
            _ = fi7
        }
    } else {
        for fi12 := 0; fi12 < 9; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 int16 = 42
                _ = len(v3)
                var v15 uint = 65535
                fmt.Println(v15, fi12)
                if true {
                    var v16 bool = true
                    v17 := GenS1{F0: 748, F1: 0}
                    if len(v2) > 0 {
                        v2[0] = v5
                    }
                    _ = v16
                    _ = v17
                } else {
                    var v18 float32 = 0.1
                    fmt.Println(v5)
                    var v19 int32 = -1
                    _ = v18
                    _ = v19
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            fmt.Println(v5)
            _ = fi12
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5)
    v5 = v1
    v6 = v6
    v6 += v6
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
