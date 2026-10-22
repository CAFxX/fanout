package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }

func main() {
    v1 := []int{9223372036854775807}
    v2 := []int{}
    var v3 int64 = 1
    v4 := []string{"foo bar", (((("a" + (("foo bar" + "a") + "foo bar")) + ("foo bar" + ((("foo bar" + ("\n\t" + "foo bar")) + "\n\t") + "\n\t"))) + "foo bar") + (("" + "hello") + "foo bar")), "hello"}
    v5 := make(map[uint32]int)
    var v6 uint32 = 535
    for fi7 := 0; fi7 < 14; fi7++ {
        _ = len(v4)
        fmt.Println(v3, v6, v3)
        if false {
            switch -45 {
            case -1:
                switch 0 {
                case 42:
                    var v8 uint32 = v6
                    _ = v8
                default:
                    for fi9 := 0; fi9 < 12; fi9++ {
                        v3 -= v3
                        var v10 int8 = 42
                        v11 := GenS1{F0: 813, F1: 249}
                        if false {
                            v2 = append(v2, fi9)
                        } else {
                            if true {
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                                var v12 bool = ((("" + (("" + "") + "")) + ((("" + "foo bar") + "foo bar") + ("a" + ("hello" + "foo bar")))) < "")
                                _ = v12
                            }
                            var v13 int64 = v3
                            var v14 int16 = -32768
                            _ = v13
                            _ = v14
                        }
                        fmt.Println(v10, fi9)
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
            case 128:
                var v15 uint32 = v6
                _ = v15
            case 773:
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 float64 = (1e10 - 3.14159)
                    _ = fi16
                    _ = v17
                }
            }
            v18 := GenS2{F0: fi7, F1: 9223372036854775807}
            var v19 int64 = v3
            _ = v18
            _ = v19
        } else {
            v3 += 853
        }
        _ = fi7
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(fi20, v3, v3)
        _ = fi20
    }
    var v21 int16 = -1
    v21 += v21
    if false {
        _ = len(v5)
    } else {
        var v22 uint64 = 18446744073709551615
        if (int(v21) >= 9223372036854775807) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v6 -= (v6 - (v6 << ((995) & 63)))
        _ = v22
    }
    var v23 int32 = 42
    if (2 > 560) {
        var v24 float64 = float64(1.5)
        _ = v24
    } else {
        v3 = (v3 * v3)
    }
    var v25 bool = ((v21 >> ((854) & 63)) == v21)
    if (!(!true)) {
        v3 = -875
        v3 = (v3 << ((65535) & 63))
        var v26 string = ""
        _ = v26
    } else {
        var v27 float64 = (1e10 / 3.14159)
        _ = v27
    }
    v21 = (^(v21))
    var v28 uint64 = 4294967295
    if v25 {
        fmt.Println(v23, v3)
        v29 := GenS1{F0: 2, F1: 110}
        v25 = true
        _ = v29
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        v31 := GenS3{F0: 93, F1: fi30, F2: 42}
        v5[919] = (fi30 % (((fi30 / (((fi30 / ((fi30) | 1))) | 1))) | 1))
        _ = fi30
        _ = v31
    }
    switch 1 {
    case -1:
        var v32 int16 = v21
        _ = v32
    case 128:
        var v33 float64 = 0.0
        _ = v33
    case 4294967295:
        v34 := GenS1{F0: 1, F1: 0}
        _ = v34
    }
    if false {
        if len(v1) > 0 {
            v1[0] = -323
        }
        var v35 int32 = v23
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
