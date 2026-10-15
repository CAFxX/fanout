package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
func genFn0() string {
    var v5 int = 32768
    var v6 int16 = -164
    v6 += -1
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 65535
    var v3 int16 = -1
    v4 := make(map[uint32]uint32)
    var v7 uint8 = 0
    v2 += (v2 & 65535)
    _ = len(v1)
    v8 := []int{0, 2147483647}
    var v9 float32 = 1.5
    switch "hello" {
    case "a":
        v10 := v4[2147483648]
        _ = v10
    case "\n\t":
        v7 -= (v7 << ((18446744073709551615) & 63))
    case "hello":
        v2 = 2
    }
    v11 := GenS1{F0: 847}
    v3 -= -1
    v3 += (^(v3))
    switch -750 {
    case 0:
        if ((42 >= 256) && (0.1 == float64(v9))) {
            fmt.Println(v9)
        } else {
            if true {
                var v12 string = (((("a" + "foo bar") + "") + "") + ("" + ("" + "\n\t")))
                _ = v12
            }
            var v13 float32 = (v9 * v9)
            if false {
                var v14 int = 945
                v15 := GenS2{F0: "\n\t"}
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v1[fi16] = false
                    _ = fi16
                }
                if (-526 < 1) {
                    if ((!(("a" + ("" + "foo bar")) < "foo bar")) && false) {
                        if ((v7 == v7) && false) {
                            for fi17 := 0; fi17 < 10; fi17++ {
                                var v18 string = "hello"
                                v2 = (v2 & v2)
                                _ = fi17
                                _ = v18
                            }
                            var v19 int16 = (v3 | (v3 >> ((715) & 63)))
                            _ = v19
                        }
                        v3 = v3
                    } else {
                        v3 += -1
                    }
                    v3 -= v3
                }
                var v20 int8 = -27
                _ = v20
            }
            _ = v13
        }
    }
    switch (("" + "hello") + "hello") {
    case "a":
        var v21 int32 = 0
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 bool = (2147483648 > 0)
        fmt.Println(fi22, v7)
        _ = fi22
        _ = v23
    }
    v24 := GenS2{F0: ""}
    v4[uint32(v7)] = 870
    var v25 int32 = -1
    v26 := GenS2{F0: ("\n\t" + "")}
    var v27 int32 = v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
