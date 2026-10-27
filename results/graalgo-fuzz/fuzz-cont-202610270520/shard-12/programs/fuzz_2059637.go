package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0(p0 uint, p1 int64, p2 uint8) int32 {
    var v4 uint64 = 1
    fmt.Println(p1, p1)
    v5 := []int{127}
    if len(v5) > 0 {
        v5[0] = 254
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        v4 = 4294967295
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1(p0 uint64, p1 bool) int8 {
    p0 = p0
    v7 := []int{-9223372036854775808, -132, 4294967295}
    v7 = append(v7, -77)
    p1 = (p1 || p1)
    _ = p0
    _ = p1
    _ = v7
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float32 = 1e10
    var v3 string = (("foo bar" + "") + "foo bar")
    var v8 int32 = 236
    fmt.Println(v3, v3)
    fmt.Println(v3, v1)
    switch "a" {
    case "\n\t":
        fmt.Println(v3, v8)
    case "a":
        switch (((((v3 + v3) + (("\n\t" + v3) + (v3 + v3))) + "\n\t") + "") + v3) {
        case "\n\t":
            v3 = (v3 + (v3 + (((v3 + "") + (v3 + (v3 + ((v3 + (("\n\t" + "") + (v3 + v3))) + v3)))) + ((v3 + v3) + v3))))
        }
    default:
        fmt.Println(v8, v8)
    }
    fmt.Println(v2)
    fmt.Println(v2)
    v9 := GenS1{F0: 255}
    for fi10 := 0; fi10 < 15; fi10++ {
        if (true && true) {
            if (127 < 1) {
                var v11 uint32 = 161
                if false {
                    if true {
                        if (127 >= 255) {
                            v12 := GenS2{F0: 32767, F1: -267, F2: 2147483647}
                            v13 := GenS1{F0: 42}
                            _ = v12
                            _ = v13
                        }
                        v14 := []int{(-(fi10)), fi10, fi10, fi10}
                        switch "foo bar" {
                        case "hello":
                            if (v11 >= (^(v11))) {
                                var v15 bool = ((v11 + 2) != (v11 % ((65535) | 1)))
                                _ = v15
                            }
                        default:
                            fmt.Println(v3)
                        }
                        _ = v14
                    }
                    v8 += v8
                    fmt.Println(v8)
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    v3 = "hello"
                    if (0 < 2) {
                        if ((v8 - v8) != v8) {
                            switch ((v3 + "\n\t") + "\n\t") {
                            case "":
                                var v17 uint = 4294967295
                                _ = v17
                            case "\n\t":
                                v18 := []int{-591, (fi16 + fi16), fi16, (^(fi16))}
                                _ = v18
                            case "foo bar":
                                var v19 uint64 = 2
                                _ = v19
                            default:
                                v3 = v3
                                v20 := GenS2{F0: 0, F1: 1, F2: (^(v11))}
                                if true {
                                    var v21 uint = 560
                                    _ = v21
                                } else {
                                    v8 -= v8
                                    var v22 int32 = -1
                                    _ = v22
                                }
                                _ = v20
                            }
                        }
                    }
                    _ = fi16
                }
                _ = v11
            } else {
                fmt.Println(fi10)
                for fi23 := 0; fi23 < 16; fi23++ {
                    fmt.Println(v1, fi23, v1)
                    v24 := GenS2{F0: -1, F1: 1000, F2: 42}
                    _ = fi23
                    _ = v24
                }
                v3 = ("" + (((v3 + ("hello" + v3)) + v3) + (v3 + ((((v3 + v3) + (("hello" + ("hello" + v3)) + ((v3 + v3) + v3))) + v3) + ((v3 + "foo bar") + v3)))))
            }
        } else {
            var v25 float32 = v2
            fmt.Println(v3, v1)
            fmt.Println(fi10, fi10, fi10)
            _ = v25
        }
        _ = fi10
    }
    v26 := GenS2{F0: 32767, F1: 1, F2: 77}
    v27 := GenS1{F0: 65535}
    v2 = v2
    fmt.Println(v3)
    for fi28 := 0; fi28 < 20; fi28++ {
        v29 := []int{-2147483648, fi28, fi28}
        fmt.Println(v1, v2, v8)
        _ = fi28
        _ = v29
    }
    var v30 bool = false
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 int16 = -916
        _ = fi31
        _ = v32
    }
    if v30 {
        fmt.Println(v2)
        v33 := []int{-186, -797, 2147483647, 32767}
        v3 = v3
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v8
    _ = v9
}
