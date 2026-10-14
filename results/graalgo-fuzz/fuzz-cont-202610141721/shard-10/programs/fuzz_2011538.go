package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int }
type GenS2 struct { F0 int8 }
func genFn0(p0 string, p1 string) uint {
    fmt.Println(p0, p1)
    var v5 uint64 = 42
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 int8, p1 int8) uint32 {
    p0 -= (p1 % ((1) | 1))
    var v6 uint32 = 65535
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (^(v6))
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]int)
    var v3 string = "a"
    v4 := make(map[uint32]bool)
    var v8 uint = genFn0("\n\t", v3)
    switch v3 {
    case "foo bar":
        switch 4294967295 {
        case -392:
            var v9 int64 = v1
            _ = v9
        }
    }
    if (1 != 12) {
        fmt.Println(v8)
        if true {
            fmt.Println(v1, v8)
            for fi10 := 0; fi10 < 3; fi10++ {
                switch (fi10 + (fi10 & 0)) {
                case 9223372036854775807:
                    v3 = ""
                default:
                    for fi11 := 0; fi11 < 19; fi11++ {
                        if (2147483647 < 2) {
                            if false {
                                if (false && true) {
                                    var v12 uint32 = 0
                                    _ = v12
                                }
                                _ = len(v4)
                                if true {
                                    _ = len(v2)
                                    v13 := v2["a"]
                                    _ = v13
                                } else {
                                    var v14 float32 = -2.25
                                    var v15 string = ("" + v3)
                                    _ = v14
                                    _ = v15
                                }
                            } else {
                                var v16 bool = ((874 <= uint64(v8)) && (!true))
                                _ = v16
                            }
                            v3 = ""
                        } else {
                            v1 = (v1 & -9223372036854775808)
                            var v17 uint8 = 55
                            _ = v17
                        }
                        fmt.Println(fi10)
                        v18 := GenS1{F0: 42, F1: 769, F2: (-(fi11))}
                        _ = fi11
                        _ = v18
                    }
                }
                for fi19 := 0; fi19 < 10; fi19++ {
                    v20 := GenS1{F0: 0, F1: -1, F2: fi19}
                    _ = fi19
                    _ = v20
                }
                _ = fi10
            }
        } else {
            var v21 int64 = v1
            v1 = v21
            _ = v21
        }
        _ = len(v4)
    }
    fmt.Println(v8)
    v8 = (v8 | 255)
    switch 32767 {
    case 127:
        if false {
            v22 := []int{1}
            _ = len(v2)
            _ = v22
        }
    }
    if (32767 != 1) {
        v8 += (v8 % ((genFn0("foo bar", v3)) | 1))
        var v23 int = 200
        v23 += v23
        _ = v23
    }
    var v24 bool = (!false)
    var v25 int = -219
    v26 := GenS1{F0: 9223372036854775807, F1: -32768, F2: v25}
    for fi27 := 0; fi27 < 10; fi27++ {
        v28 := GenS2{F0: int8(v8)}
        for fi29 := 0; fi29 < 17; fi29++ {
            var v30 string = v3
            v24 = ((v30 + (v30 + v3)) <= v3)
            v31 := GenS2{F0: 42}
            v3 = "foo bar"
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v8
}
