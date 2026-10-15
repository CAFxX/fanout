package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 string }
func genFn0(p0 uint32) int {
    var v7 uint64 = 4294967295
    var v8 uint8 = 0
    switch -9223372036854775808 {
    case 42:
        v9 := GenS1{F0: "foo bar", F1: (p0 << ((18446744073709551615) & 63))}
        _ = v9
    }
    var v10 int = 128
    v11 := []int{v10, 0}
    _ = p0
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    return (v10 & v10)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{4294967295, 2, 1, 1}
    var v3 bool = (0 == 18446744073709551615)
    v4 := []int{}
    var v5 int64 = 42
    var v6 uint64 = 992
    v12 := GenS1{F0: ((("\n\t" + "a") + "foo bar") + "foo bar"), F1: 2147483647}
    if v3 {
        var v13 int64 = v5
        _ = v13
    }
    var v14 bool = v3
    var v15 uint16 = 131
    switch "a" {
    case "":
        if (-244 > 138) {
            var v16 string = ("" + "foo bar")
            if (42 > -1) {
                if v14 {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        if true {
                            v5 = v5
                        } else {
                            fmt.Println(v6)
                            var v18 uint = 42
                            _ = v18
                        }
                        var v19 uint32 = 312
                        var v20 int32 = 0
                        _ = fi17
                        _ = v19
                        _ = v20
                    }
                }
                var v21 uint64 = v6
                _ = v21
            } else {
                for fi22 := 0; fi22 < 9; fi22++ {
                    v3 = v14
                    fmt.Println(v6)
                    _ = fi22
                }
                switch v16 {
                case "foo bar":
                    v23 := GenS2{F0: v6, F1: 0, F2: (v16 + v16)}
                    _ = v23
                case "":
                    if len(v4) > 0 {
                        v4[0] = 32768
                    }
                default:
                    var v24 string = v16
                    _ = v24
                }
            }
            _ = v16
        } else {
            for fi25 := 0; fi25 < 18; fi25++ {
                var v26 float32 = ((0.0 * 0.1) - 0.1)
                _ = fi25
                _ = v26
            }
        }
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    _ = len(v1)
    if (2147483647 < 2147483647) {
        if (((!(((v15 / (((v15 & v15)) | 1)) >= v15) && (!v3))) || (!(268 > 1))) && (v3 && true)) {
            if v3 {
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v14 = true
                }
            }
            v27 := GenS1{F0: (("a" + "\n\t") + ("foo bar" + ("foo bar" + "\n\t"))), F1: 33}
            if ((^(v6)) != v6) {
                v3 = false
                v5 = 936
                for fi28 := 0; fi28 < 19; fi28++ {
                    v29 := GenS2{F0: v6, F1: genFn0(943), F2: "\n\t"}
                    var v30 string = ("a" + "foo bar")
                    fmt.Println(v14, v3, v5)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            } else {
                v6 += 18446744073709551615
                var v31 uint8 = 42
                _ = v31
            }
            _ = v27
        }
        v32 := GenS1{F0: ("\n\t" + "hello"), F1: uint32(v15)}
        var v33 uint64 = (v6 >> ((42) & 63))
        _ = v32
        _ = v33
    } else {
        fmt.Println(v3)
        v6 = v6
        var v34 int8 = -1
        _ = v34
    }
    v4 = append(v4, 0)
    if (503 > 591) {
        v15 -= v15
        v14 = (942 <= -390)
        var v35 string = "\n\t"
        _ = v35
    }
    var v36 uint16 = (v15 & 65535)
    var v37 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
