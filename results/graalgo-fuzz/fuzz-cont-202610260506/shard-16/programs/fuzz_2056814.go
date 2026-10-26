package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint64; F2 int32 }
func genFn0(p0 uint32, p1 int) uint32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 string = "foo bar"
        _ = fi5
        _ = v6
    }
    switch p1 {
    case -9223372036854775808:
        p0 = 1
    case 32768:
        p0 -= (p0 % ((p0) | 1))
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 bool) int8 {
    p0 = (3.14159 == (0.0 / 0.1))
    v7 := GenS1{F0: 18446744073709551615, F1: 100, F2: -1}
    _ = p0
    _ = v7
    return 0
}

func main() {
    var v1 int64 = -961
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    fmt.Println(v1)
    var v8 bool = (false && true)
    var v9 int16 = -849
    for fi10 := 0; fi10 < 6; fi10++ {
        if ((genFn0(246, (fi10 % (((^(fi10))) | 1))) < 2) || ((((1e10 + 0.0) / 3.14159) - (1.5 + 1e10)) <= (1e10 / 1.0))) {
            if false {
                switch ("\n\t" + "hello") {
                case "a":
                    for fi11 := 0; fi11 < 4; fi11++ {
                        v1 -= 826
                        switch fi11 {
                        case 32768:
                            v9 += -1
                        case 1:
                            var v12 bool = v8
                            _ = v12
                        }
                        fmt.Println(v8, v9)
                        _ = fi11
                    }
                default:
                    v4[65535] = "foo bar"
                }
                v13 := GenS2{F0: -1}
                v14 := []int{(^(fi10)), -2147483648, (fi10 + (fi10 >> ((345) & 63))), (-(fi10))}
                _ = v13
                _ = v14
            } else {
                v15 := GenS2{F0: -988}
                var v16 float64 = 1.5
                _ = v15
                _ = v16
            }
        }
        v1 += int64(v9)
        _ = fi10
    }
    var v17 int8 = 0
    if (!v8) {
        var v18 int = 2147483647
        switch (^(v18)) {
        case 2147483647:
            if v8 {
                v1 = (v1 + v1)
            } else {
                for fi19 := 0; fi19 < 7; fi19++ {
                    v9 += (v9 * (v9 / ((v9) | 1)))
                    v4[2147483647] = "foo bar"
                    _ = fi19
                }
            }
        case 9223372036854775807:
            v17 += (v17 ^ genFn1((!(true && (v8 && true)))))
        case 32768:
            var v20 float32 = ((((((((1e10 + 1e10) / 3.14159) / (1e10 * 1e10)) / -2.25) + 0.1) + (((1.5 - 1e10) - 0.0) + (1e10 - -2.25))) + (-2.25 - ((0.1 / 1.5) - float32(3.14159)))) - 3.14159)
            _ = v20
        }
        v9 = 42
        _ = v18
    } else {
        v17 = v17
    }
    var v21 int64 = (v1 ^ (^(v1)))
    var v22 float32 = 1e10
    v17 -= v17
    var v23 int64 = (v21 % ((-1) | 1))
    var v24 int = 1
    v25 := GenS3{F0: ((("foo bar" + (("\n\t" + ("hello" + ("\n\t" + "\n\t"))) + "\n\t")) + ("" + "foo bar")) + "a"), F1: 646, F2: int32(v24)}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
