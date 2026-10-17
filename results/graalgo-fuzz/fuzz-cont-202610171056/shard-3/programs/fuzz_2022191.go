package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0(p0 int32) int16 {
    var v5 int8 = 0
    if true {
        if false {
            for fi6 := 0; fi6 < 7; fi6++ {
                var v7 int = fi6
                _ = fi6
                _ = v7
            }
        }
    }
    p0 = p0
    v8 := []int{32767, -292, 32767}
    _ = p0
    _ = v5
    _ = v8
    return 1
}
func genFn1() uint64 {
    var v9 uint64 = 663
    v9 -= v9
    v10 := GenS1{F0: uint(v9)}
    if (-128 <= -128) {
        if ((!false) && (1 <= 835)) {
            fmt.Println(v9)
            if ((4294967295 == 2147483647) || true) {
                for fi11 := 0; fi11 < 5; fi11++ {
                    v9 = (v9 << ((830) & 63))
                    _ = fi11
                }
                var v12 int32 = 538
                _ = v12
            } else {
                switch -224 {
                case 127:
                    v13 := []int{-524, -9223372036854775808}
                    _ = v13
                case -306:
                    var v14 string = "a"
                    _ = v14
                default:
                    v15 := GenS1{F0: 65535}
                    _ = v15
                }
                switch ((("a" + ("a" + "")) + "\n\t") + "\n\t") {
                case "hello":
                    for fi16 := 0; fi16 < 15; fi16++ {
                        v17 := []int{}
                        _ = fi16
                        _ = v17
                    }
                case "":
                    fmt.Println(v9)
                case "foo bar":
                    v18 := GenS2{F0: 42, F1: -2147483648, F2: 0}
                    _ = v18
                default:
                    if false {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            if false {
                                v20 := []int{}
                                _ = v20
                            }
                            _ = fi19
                        }
                    }
                }
                v21 := []int{32768, -1, -533, -9223372036854775808}
                _ = v21
            }
            switch "\n\t" {
            case "foo bar":
                v22 := []int{-9223372036854775808}
                _ = v22
            }
        }
        v9 -= v9
    } else {
        v9 = v9
        for fi23 := 0; fi23 < 14; fi23++ {
            v9 -= (v9 ^ v9)
            v24 := GenS1{F0: 1}
            _ = fi23
            _ = v24
        }
    }
    _ = v10
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int64 = 42
    v2 := []uint32{4294967295, 2147483648}
    var v3 int64 = 0
    v4 := []uint32{367}
    v4 = append(v4, 546)
    switch -579 {
    case 0:
        if (42 <= -32768) {
            fmt.Println(v3, v1)
        } else {
            v3 -= v1
            v1 += 132
            v3 = v1
        }
    case 32767:
        switch (("\n\t" + "a") + ("foo bar" + (("a" + "\n\t") + ("hello" + "foo bar")))) {
        case "":
            var v25 int32 = 42
            _ = v25
        }
    default:
        v26 := GenS2{F0: 615, F1: 459, F2: 343}
        _ = v26
    }
    var v27 int8 = -128
    v1 += -9223372036854775808
    switch 9223372036854775807 {
    case 2147483647:
        for fi28 := 0; fi28 < 12; fi28++ {
            v27 -= (v27 / ((v27) | 1))
            _ = fi28
        }
    case 128:
        var v29 uint64 = 4294967295
        _ = v29
    case -941:
        var v30 uint32 = 2
        _ = v30
    default:
        var v31 int32 = int32(v27)
        _ = v31
    }
    v1 += (-(v1))
    v27 -= v27
    var v32 int32 = -13
    if ("hello" > "") {
        v3 = (v3 * 921)
        var v33 uint64 = 42
        _ = v33
    }
    switch "" {
    case "":
        var v34 uint64 = 2
        _ = v34
    default:
        if true {
            v4 = append(v4, 663)
            var v35 uint32 = 4294967295
            _ = v35
        }
    }
    v1 -= v3
    var v36 int32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v36
    _ = v4
}
