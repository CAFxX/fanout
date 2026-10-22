package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int16 }
func genFn0(p0 string, p1 uint16, p2 int64) uint {
    switch 32768 {
    case 0:
        p0 = p0
    case -739:
        var v5 uint = 18446744073709551615
        _ = v5
    case 127:
        fmt.Println(p2, p2)
    }
    if ("\n\t" <= p0) {
        var v6 int64 = (p2 & p2)
        var v7 uint8 = 80
        _ = v6
        _ = v7
    }
    var v8 int32 = -101
    var v9 uint = 520
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return v9
}

func main() {
    v1 := []uint32{2147483648, 980, 844}
    var v2 uint = 292
    v3 := []string{(("foo bar" + "") + "a")}
    var v4 bool = (!(1.5 > 0.1))
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    fmt.Println(v4, v4)
    if false {
        v2 = (v2 + 42)
        switch 0 {
        case 2147483647:
            var v10 int32 = 2147483647
            _ = v10
        case 32767:
            v4 = v4
        default:
            if true {
                var v11 int32 = -1
                for fi12 := 0; fi12 < 5; fi12++ {
                    v13 := GenS1{F0: (v2 << (((v2 + (v2 << ((genFn0("foo bar", 1, 42)) & 63)))) & 63)), F1: 0, F2: 127}
                    v14 := GenS1{F0: v2, F1: 2, F2: 224}
                    if (int32(fi12) > 1) {
                        if (v4 || (("" + "\n\t") < "")) {
                            fmt.Println(v11)
                            var v15 float32 = 1e10
                            var v16 uint32 = 65535
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS1{F0: v2, F1: 65535, F2: 42}
                        _ = v17
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = v11
            } else {
                v2 = (v2 + v2)
            }
        }
        switch ("\n\t" + "") {
        case "":
            var v18 uint8 = 128
            _ = v18
        case "a":
            var v19 string = "\n\t"
            _ = v19
        }
    }
    var v20 float64 = 1.5
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = append(v1, 208)
        v20 += v20
    } else {
        v21 := GenS1{F0: (^(v2)), F1: 2, F2: 1}
        _ = v21
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    fmt.Println(v4, v2, v4)
    v22 := GenS1{F0: v2, F1: 2, F2: 58}
    if ("a" != "hello") {
        v23 := GenS2{F0: 2147483647, F1: 100, F2: 32767}
        v2 = 317
        _ = v23
    }
    v4 = v4
    switch 127 {
    case -92:
        v24 := GenS1{F0: (v2 ^ genFn0("a", 724, 429)), F1: 42, F2: 128}
        _ = v24
    case 127:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    if (v4 || (4294967295 == 4294967295)) {
        v25 := GenS2{F0: 2147483647, F1: 100, F2: -1}
        if len(v3) > 0 {
            v3[0] = ("\n\t" + "")
        }
        var v26 uint32 = 42
        _ = v25
        _ = v26
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
}
