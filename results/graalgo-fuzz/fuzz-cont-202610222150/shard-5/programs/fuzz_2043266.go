package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
func genFn0(p0 string) uint32 {
    p0 = ("a" + p0)
    if (0 < -1) {
        p0 = ((p0 + p0) + "foo bar")
    } else {
        var v4 int8 = 127
        v5 := []int{-677, 4294967295, 436, 9223372036854775807}
        fmt.Println(p0)
        _ = v4
        _ = v5
    }
    var v6 float64 = -2.25
    _ = p0
    _ = v6
    return 2
}

func main() {
    var v1 uint64 = 79
    var v2 bool = true
    v3 := make(map[int]int)
    var v7 int = -9223372036854775808
    v1 = (v1 >> ((18446744073709551615) & 63))
    var v8 int16 = 0
    v9 := GenS1{F0: "hello"}
    fmt.Println(v8, v7)
    var v10 float64 = 0.0
    var v11 int64 = -9223372036854775808
    v11 += v11
    if false {
        var v12 int64 = v11
        if (1e10 > 3.14159) {
            switch ("hello" + ("a" + "\n\t")) {
            case "":
                if v2 {
                    v13 := []int{4294967295, v7}
                    v7 = 42
                    var v14 bool = v2
                    _ = v13
                    _ = v14
                } else {
                    var v15 float64 = ((3.14159 * (v10 * ((float64(1e10) - 1.5) * 3.14159))) - 0.0)
                    v12 = -654
                    _ = v15
                }
            default:
                switch ("hello" + "") {
                case "":
                    v16 := GenS2{F0: (^(v7))}
                    _ = v16
                case "hello":
                    for fi17 := 0; fi17 < 3; fi17++ {
                        switch (^(fi17)) {
                        case 42:
                            v1 += 566
                        case 32767:
                            if (0 > 299) {
                                if v2 {
                                    v10 = v10
                                    fmt.Println(v8, v12)
                                }
                            } else {
                                v1 = 2
                                v18 := GenS2{F0: 32767}
                                _ = v18
                            }
                        }
                        _ = fi17
                    }
                case "a":
                    v10 = v10
                default:
                    v19 := GenS2{F0: v7}
                    _ = v19
                }
            }
            v20 := GenS2{F0: v7}
            var v21 uint32 = 42
            _ = v20
            _ = v21
        } else {
            fmt.Println(v2, v11, v2)
            _ = len(v3)
            var v22 bool = (((3.14159 * 1.5) > (0.1 * (1.5 - 1.5))) || (!v2))
            _ = v22
        }
        var v23 uint8 = 47
        _ = v12
        _ = v23
    }
    v24 := []int{32767}
    v8 -= v8
    v11 = (v11 ^ (v11 - v11))
    switch -9223372036854775808 {
    case 0:
        for fi25 := 0; fi25 < 16; fi25++ {
            fmt.Println(fi25)
            var v26 uint8 = 184
            for fi27 := 0; fi27 < 17; fi27++ {
                var v28 uint64 = (^(v1))
                var v29 uint32 = 873
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi25
            _ = v26
        }
    }
    v30 := GenS2{F0: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v7
    _ = v8
    _ = v9
}
