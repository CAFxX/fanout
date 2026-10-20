package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 int32; F1 int16; F2 uint32 }
type GenS3 struct { F0 int8; F1 int64; F2 int64 }
type GenS4 struct { F0 int; F1 int16; F2 string }
func genFn0() uint64 {
    if false {
        if (-1 != -808) {
            v4 := []int{2147483647, 2147483647, -754}
            var v5 bool = (true && (-2147483648 > -184))
            _ = v4
            _ = v5
        }
    }
    var v6 uint32 = 480
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(fi7, v6)
            if true {
                if false {
                    v9 := []int{(-(fi8)), fi8}
                    if (fi7 <= (fi7 + (fi8 | (fi8 ^ (-(fi8)))))) {
                        switch (fi8 >> ((205) & 63)) {
                        case 4294967295:
                            if true {
                                if true {
                                    _ = len(v9)
                                    var v10 int8 = -31
                                    _ = len(v9)
                                    _ = v10
                                } else {
                                    var v11 int64 = 9223372036854775807
                                    v12 := GenS1{F0: 0, F1: (false && (!(("hello" + "hello") >= "foo bar"))), F2: true}
                                    v6 += v6
                                    switch ((("" + "a") + "hello") + ("" + ((("\n\t" + (("a" + "foo bar") + "foo bar")) + ("hello" + "a")) + "\n\t"))) {
                                    case "a":
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    default:
                                        v6 = (v6 / ((65535) | 1))
                                    }
                                    v13 := GenS2{F0: 0, F1: 32767, F2: v6}
                                    _ = v11
                                    _ = v12
                                    _ = v13
                                }
                                fmt.Println(fi8, fi7, fi7)
                                v6 = (v6 | 65535)
                                v9 = append(v9, (-(fi7)))
                            }
                        default:
                            v6 = v6
                        }
                        if len(v9) > 0 {
                            v9[0] = fi7
                        }
                    }
                    var v14 int32 = -1
                    _ = v14
                    _ = v9
                } else {
                    v15 := []int{fi7}
                    fmt.Println(fi8)
                    var v16 uint8 = 128
                    _ = v15
                    _ = v16
                }
                v17 := GenS3{F0: 127, F1: 9223372036854775807, F2: 1}
                v6 += 65535
                var v18 int16 = -1
                _ = v17
                _ = v18
            } else {
                if true {
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        v19 := []int{(fi8 - (fi8 * fi7)), fi8, fi7, (fi8 >> ((0) & 63))}
                        _ = v19
                    case "\n\t":
                        v6 = (^(v6))
                    default:
                        switch ("" + ("a" + "")) {
                        case "hello":
                            var v20 int = (fi8 << ((65535) & 63))
                            _ = v20
                        case "":
                            var v21 int8 = -87
                            _ = v21
                        }
                    }
                    switch "" {
                    case "\n\t":
                        var v22 uint32 = (v6 ^ 279)
                        _ = v22
                    }
                    var v23 uint8 = 42
                    _ = v23
                }
                if false {
                    var v24 uint8 = 1
                    _ = v24
                }
            }
            v6 = (v6 & 0)
            _ = fi8
        }
        v25 := GenS1{F0: 127, F1: (float64(1e10) > ((0.0 / 1.0) + (float64(0.0) - 1e10))), F2: true}
        v6 += (v6 ^ v6)
        _ = fi7
        _ = v25
    }
    if ((42 == 260) || true) {
        var v26 int32 = -980
        _ = v26
    }
    _ = v6
    return 257
}
func genFn1() uint64 {
    if (false && (!(false && false))) {
        var v27 bool = true
        fmt.Println(v27)
        v27 = ((v27 || true) && (!v27))
        _ = v27
    } else {
        v28 := []int{1}
        _ = len(v28)
        _ = v28
    }
    var v29 int16 = 1000
    _ = v29
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 2
    var v3 int32 = -649
    fmt.Println(v3)
    var v30 int16 = 32767
    for fi31 := 0; fi31 < 3; fi31++ {
        switch (fi31 & fi31) {
        case 32767:
            var v32 uint = 740
            _ = v32
        default:
            if (v3 != v3) {
                for fi33 := 0; fi33 < 19; fi33++ {
                    var v34 uint64 = 0
                    v35 := GenS4{F0: 60, F1: 650, F2: "a"}
                    v36 := GenS3{F0: 41, F1: -9223372036854775808, F2: 167}
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
            }
        }
        v3 = (^(v3))
        fmt.Println(v3, v3, fi31)
        _ = fi31
    }
    fmt.Println(v2, v30, v3)
    var v37 int16 = (v30 ^ int16(v3))
    var v38 uint64 = 4294967295
    v37 += v30
    fmt.Println(v30)
    if false {
        var v39 uint32 = 2147483648
        switch "" {
        case "":
            v37 -= 32767
        }
        _ = v39
    }
    if true {
        v37 = v37
        v37 = v37
    }
    delete(v1, 4294967295)
    var v40 int16 = int16(v38)
    var v41 string = ((("a" + "") + "\n\t") + ("a" + ("a" + ("a" + ""))))
    v41 = v41
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v40
    _ = v41
}
