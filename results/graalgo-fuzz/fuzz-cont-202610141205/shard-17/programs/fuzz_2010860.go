package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint8; F1 int8; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 int64; F1 string; F2 int16 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 bool) int16 {
    if p0 {
        fmt.Println(p0)
        var v6 int = -813
        fmt.Println(p0, v6)
        _ = v6
    }
    var v7 string = "a"
    var v8 uint32 = 119
    _ = p0
    _ = v7
    _ = v8
    return -491
}
func genFn1(p0 int64, p1 uint, p2 uint32) uint {
    var v9 int8 = -128
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{-9223372036854775808, fi10, fi10, fi10}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p1 << ((2) & 63))
}

func main() {
    var v1 float32 = 0.1
    var v2 uint = 0
    v3 := make(map[string]int)
    var v4 string = ("foo bar" + ("hello" + ("hello" + "foo bar")))
    v5 := make(map[string]uint32)
    var v12 uint64 = 0
    v13 := GenS1{F0: 35, F1: -18}
    var v14 uint8 = 141
    switch (v4 + "a") {
    case "a":
        var v15 int8 = -128
        _ = v15
    }
    var v16 int = -2147483648
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 string = v4
        var v19 int32 = 0
        v19 += (-(v19))
        _ = fi17
        _ = v18
        _ = v19
    }
    if ((v4 >= (v4 + "foo bar")) || (!(4294967295 < uint32(v2)))) {
        var v20 uint16 = 42
        _ = v20
    }
    fmt.Println(v1, v2)
    v4 = (("a" + v4) + (((("hello" + v4) + v4) + (v4 + ("a" + ((v4 + "hello") + "\n\t")))) + ("a" + (v4 + v4))))
    v21 := GenS1{F0: 2, F1: -1}
    for fi22 := 0; fi22 < 8; fi22++ {
        if (true || true) {
            fmt.Println(fi22)
            if false {
                v23 := []int{fi22, (v16 >> ((18446744073709551615) & 63))}
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                switch (v4 + (v4 + "a")) {
                case "a":
                    _ = len(v3)
                case "hello":
                    var v24 uint32 = uint32(v12)
                    _ = v24
                }
                _ = v23
            }
            v25 := GenS2{F0: 216, F1: 1, F2: int16(v2)}
            _ = v25
        }
        _ = fi22
    }
    var v26 string = (v4 + v4)
    if false {
        fmt.Println(v26, v1)
        v26 = v4
        if true {
            v27 := GenS2{F0: v14, F1: 1, F2: -289}
            if (v16 < int(v2)) {
                var v28 bool = (((^(v12)) < (^(v12))) && (605 < 1))
                _ = v28
            } else {
                var v29 int32 = 0
                _ = v29
            }
            switch (^(v16)) {
            case -2147483648:
                switch ("a" + ("hello" + "")) {
                case "hello":
                    var v30 float64 = (-2.25 / 1e10)
                    _ = v30
                case "a":
                    v31 := GenS2{F0: 32, F1: 0, F2: -640}
                    _ = v31
                }
            }
            _ = v27
        } else {
            if true {
                switch v16 {
                case 628:
                    if false {
                        delete(v5, "")
                        switch "" {
                        case "hello":
                            switch "" {
                            case "a":
                                v16 = (-(v16))
                            case "":
                                if (97 >= 2) {
                                    v32 := GenS3{F0: ((v14 - 11) == v14)}
                                    v2 += 0
                                    fmt.Println(v2, v1)
                                    _ = v32
                                } else {
                                    v33 := GenS4{F0: 42, F1: 0, F2: 2}
                                    _ = v33
                                }
                            case "hello":
                                v26 = v26
                            }
                        default:
                            v34 := GenS5{F0: v2, F1: (v12 | 0)}
                            _ = v34
                        }
                        v35 := GenS6{F0: 42, F1: (v4 + "foo bar"), F2: genFn0(false)}
                        _ = v35
                    }
                default:
                    var v36 bool = (189 <= v14)
                    _ = v36
                }
                v1 += 3.14159
            } else {
                fmt.Println(v26)
                v1 += (0.0 - v1)
            }
        }
    } else {
        var v37 string = "hello"
        if true {
            if (false || true) {
                v16 = v16
            } else {
                var v38 uint64 = (v12 % ((v12) | 1))
                fmt.Println(v14)
                _ = v38
            }
            v39 := GenS1{F0: 4294967295, F1: -128}
            _ = v39
        } else {
            v40 := GenS6{F0: -9223372036854775808, F1: v4, F2: -32768}
            _ = v40
        }
        _ = v37
    }
    v41 := GenS7{F0: 255}
    var v42 uint = v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
