package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint16, p2 bool) int16 {
    p2 = p2
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            switch fi4 {
            case -2147483648:
                fmt.Println(fi4)
            case 300:
                var v6 uint8 = 184
                _ = v6
            default:
                if true {
                    if (fi4 == (fi4 % ((-9223372036854775808) | 1))) {
                        v7 := []int{(^(fi4)), -9223372036854775808, 0, (fi5 / ((fi5) | 1))}
                        p1 += (p1 + (p1 * 1))
                        v8 := GenS1{F0: 2, F1: 2147483647}
                        _ = v7
                        _ = v8
                    } else {
                        var v9 uint8 = 255
                        v9 = v9
                        _ = v9
                    }
                    var v10 int16 = 32767
                    _ = v10
                } else {
                    var v11 string = ("\n\t" + "\n\t")
                    v12 := GenS1{F0: (p1 % ((p1) | 1)), F1: -412}
                    var v13 int = fi4
                    _ = v11
                    _ = v12
                    _ = v13
                }
            }
            v14 := GenS1{F0: p1, F1: int32(fi4)}
            v15 := GenS2{F0: 213, F1: 4294967295, F2: p1}
            _ = fi5
            _ = v14
            _ = v15
        }
        _ = fi4
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]string)
    var v3 int = 0
    v16 := []int{(v3 + (-(v3))), v3}
    if (!(true && true)) {
        v3 = (v3 / ((-9223372036854775808) | 1))
        if true {
            v1 = (v1 + v1)
            for fi17 := 0; fi17 < 8; fi17++ {
                fmt.Println(v3, v3, v1)
                var v18 uint16 = 0
                _ = fi17
                _ = v18
            }
            v3 = (v3 + v3)
        }
    } else {
        v3 = 32768
        fmt.Println(v1, v3)
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 uint32 = 0
        v2[v3] = "foo bar"
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: 0, F1: 923}
    if (-2147483648 != -2147483648) {
        v3 -= -1
    }
    v22 := v2[(v3 ^ (-(v3)))]
    v23 := GenS3{F0: 540, F1: 972}
    v24 := GenS2{F0: 2, F1: 217, F2: 65535}
    v22 = (v22 + "\n\t")
    delete(v2, -9223372036854775808)
    if false {
        if false {
            var v25 int32 = 0
            switch "foo bar" {
            case "":
                if false {
                    fmt.Println(v22, v3, v25)
                    v22 = "\n\t"
                } else {
                    fmt.Println(v1)
                    v22 = ((("" + v22) + v22) + v22)
                    switch v22 {
                    case "":
                        v26 := v2[4294967295]
                        _ = v26
                    case "a":
                        var v27 uint = 182
                        _ = v27
                    }
                }
            case "hello":
                v28 := GenS4{F0: 514, F1: 0}
                _ = v28
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                var v30 int8 = 127
                _ = fi29
                _ = v30
            }
            _ = v25
        } else {
            v1 = float32(((1e10 * 1.5) - 1e10))
            switch "a" {
            case "hello":
                v22 = ("" + "\n\t")
            case "a":
                _ = len(v16)
            default:
                v3 += 32768
            }
            v31 := GenS5{F0: 4294967295}
            _ = v31
        }
        v3 = 1
        var v32 uint16 = 65535
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 int16 = genFn0(-1, 1, (508 < 0))
        var v35 int64 = 213
        v34 += (v34 | genFn0(-108, 2, ((v22 + "") < v22)))
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
