package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0() int16 {
    var v6 int64 = -9223372036854775808
    var v7 int = 4294967295
    var v8 int32 = 701
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{4294967295, 2147483648}
    var v3 int32 = 671
    var v4 uint32 = 65535
    var v5 string = ("" + "a")
    switch 127 {
    case -330:
        fmt.Println(v3, v3, v4)
    default:
        if true {
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 int8 = 127
                _ = fi9
                _ = v10
            }
            v3 = v3
        } else {
            var v11 uint8 = 255
            v12 := GenS1{F0: -32768, F1: "foo bar"}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v14 uint = 0
        _ = fi13
        _ = v14
    }
    switch 1 {
    case 714:
        v4 -= v4
    case -166:
        v2 = append(v2, (v4 - v4))
    case 899:
        var v15 uint32 = (v4 + 755)
        _ = v15
    default:
        var v16 uint32 = 42
        _ = v16
    }
    var v17 float32 = (((3.14159 * 0.1) / -2.25) - (3.14159 * 1.5))
    v18 := GenS1{F0: 1, F1: v5}
    v4 = v4
    v2 = append(v2, (v4 ^ (v4 | 1)))
    var v19 int16 = genFn0()
    if false {
        switch 156 {
        case 4294967295:
            if (false && true) {
                var v20 uint32 = 65535
                if (v19 != -1) {
                    if false {
                        switch (v5 + v5) {
                        case "foo bar":
                            var v21 int32 = 1
                            _ = v21
                        case "hello":
                            var v22 string = (v5 + (("hello" + v5) + "foo bar"))
                            _ = v22
                        default:
                            var v23 float32 = v17
                            _ = v23
                        }
                        v19 -= (v19 + -32768)
                    }
                    var v24 int16 = -5
                    var v25 bool = true
                    _ = v24
                    _ = v25
                }
                _ = v20
            }
        case 1:
            var v26 uint8 = 128
            _ = v26
        default:
            var v27 uint32 = (v4 * v4)
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            v1[v5] = v5
            v29 := GenS2{F0: 287, F1: 0}
            if (true && false) {
                if (uint8(v4) >= 128) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v19)
                    switch (v5 + (v5 + "hello")) {
                    case "\n\t":
                        v30 := GenS1{F0: v19, F1: "\n\t"}
                        _ = v30
                    }
                } else {
                    if ((v19 >> ((18446744073709551615) & 63)) < (^(v19))) {
                        if len(v2) > 0 {
                            v2[0] = v4
                        }
                    }
                    var v31 uint16 = 32767
                    _ = v31
                }
                if false {
                    v32 := GenS1{F0: 0, F1: v5}
                    v33 := v1["hello"]
                    _ = v32
                    _ = v33
                } else {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        v35 := GenS1{F0: v19, F1: v5}
                        var v36 float32 = 0.1
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
            }
            _ = fi28
            _ = v29
        }
        switch 32768 {
        case -705:
            fmt.Println(v19, v3)
        case 128:
            fmt.Println(v3, v17)
        }
        v37 := GenS2{F0: 820, F1: 18446744073709551615}
        _ = v37
    }
    v19 = -32768
    v3 += (v3 ^ v3)
    if ((false && true) || true) {
        fmt.Println(v17, v17, v17)
        _ = len(v2)
    } else {
        _ = len(v2)
        var v38 uint32 = (v4 - (^(v4)))
        var v39 bool = (true || true)
        _ = v38
        _ = v39
    }
    delete(v1, v5)
    if false {
        v40 := v1["foo bar"]
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
