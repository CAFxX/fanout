package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := []string{"\n\t", (("\n\t" + "") + ("a" + "hello"))}
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 uint = 42
    v5 := []string{"hello", "foo bar", "a", ("" + "foo bar")}
    var v6 uint = 255
    fmt.Println(v6)
    v4 -= (v6 << (((v6 + (v4 * (v4 | v6)))) & 63))
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 128, F1: (("" + "hello") + ("hello" + "")), F2: 2}
        _ = fi7
        _ = v8
    }
    v4 += (v6 >> (((^(v4))) & 63))
    var v9 float32 = 0.0
    for fi10 := 0; fi10 < 13; fi10++ {
        v6 += 0
        _ = fi10
    }
    if (!((1 != -36) && true)) {
        for fi11 := 0; fi11 < 4; fi11++ {
            switch (("a" + "") + "\n\t") {
            case "":
                for fi12 := 0; fi12 < 8; fi12++ {
                    v6 -= 2
                    _ = fi12
                }
            case "a":
                var v13 string = "\n\t"
                _ = v13
            }
            var v14 int16 = 42
            var v15 int64 = -1
            _ = fi11
            _ = v14
            _ = v15
        }
        _ = len(v3)
    } else {
        var v16 uint = (v4 * (v6 >> (((v6 | (v4 - v4))) & 63)))
        switch 32767 {
        case -9223372036854775808:
            v17 := GenS2{F0: 1}
            _ = v17
        case -226:
            v18 := GenS3{F0: 4294967295, F1: 42, F2: 100}
            _ = v18
        case 0:
            for fi19 := 0; fi19 < 2; fi19++ {
                var v20 bool = ((fi19 & (fi19 - (fi19 & (fi19 << ((v16) & 63))))) <= fi19)
                for fi21 := 0; fi21 < 1; fi21++ {
                    var v22 uint16 = 1
                    switch -764 {
                    case 42:
                        v23 := GenS3{F0: 488, F1: 184, F2: -1}
                        _ = v23
                    case 4294967295:
                        var v24 uint = uint(v22)
                        _ = v24
                    }
                    v25 := v2[4294967295]
                    v6 = v4
                    _ = fi21
                    _ = v22
                    _ = v25
                }
                _ = fi19
                _ = v20
            }
        default:
            switch -673 {
            case 2147483647:
                for fi26 := 0; fi26 < 18; fi26++ {
                    v27 := GenS4{F0: 343}
                    _ = fi26
                    _ = v27
                }
            case 658:
                v6 += (v16 - (v4 ^ v16))
            default:
                for fi28 := 0; fi28 < 1; fi28++ {
                    if (2147483647 >= 2147483647) {
                        fmt.Println(v16, v16, v9)
                        for fi29 := 0; fi29 < 18; fi29++ {
                            v30 := v2[42]
                            var v31 string = ""
                            fmt.Println(v9, v6)
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        v3["foo bar"] = "hello"
                    }
                    v6 = 18446744073709551615
                    _ = fi28
                }
            }
        }
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
