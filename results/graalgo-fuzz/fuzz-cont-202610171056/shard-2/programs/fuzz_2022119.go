package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 uint64 }
func genFn0() int64 {
    var v6 int16 = 648
    v7 := []int{0, -9223372036854775808, 795}
    var v8 int = 32767
    _ = v6
    _ = v7
    _ = v8
    return 681
}

func main() {
    var v1 uint64 = 1
    v2 := []int{1, -419, 746}
    v3 := []uint32{}
    var v4 uint16 = 92
    v5 := make(map[uint32]int)
    var v9 uint8 = 0
    if false {
        fmt.Println(v4, v9)
    }
    switch "\n\t" {
    case "hello":
        if (!false) {
            fmt.Println(v4, v9)
            v9 = 103
        }
    }
    v4 = (v4 & (v4 << ((65535) & 63)))
    v10 := v5[0]
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            v9 = 4
            var v13 float32 = 0.1
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 18446744073709551615, F1: 2}
        for fi15 := 0; fi15 < 15; fi15++ {
            v1 -= 0
            if false {
                fmt.Println(v4, v4, v1)
                if true {
                    var v16 int = (fi15 ^ (v10 * (fi11 >> ((4294967295) & 63))))
                    v1 += uint64(v4)
                    if true {
                        v17 := GenS1{F0: (v1 << ((4294967295) & 63)), F1: 1}
                        fmt.Println(fi15, fi11)
                        var v18 bool = false
                        _ = v17
                        _ = v18
                    }
                    _ = v16
                } else {
                    switch "" {
                    case "":
                        if false {
                            if false {
                                var v19 uint8 = v9
                                var v20 uint64 = v1
                                _ = v19
                                _ = v20
                            }
                        } else {
                            var v21 string = ("foo bar" + "a")
                            var v22 int64 = -980
                            v1 -= v1
                            var v23 int64 = (v22 ^ (v22 % ((0) | 1)))
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    case "foo bar":
                        var v24 uint64 = v1
                        _ = v24
                    case "a":
                        var v25 uint8 = v9
                        _ = v25
                    }
                    switch ("foo bar" + ("" + ((("foo bar" + "") + "foo bar") + (("foo bar" + (((("\n\t" + ("hello" + ("hello" + "hello"))) + ("" + "a")) + "hello") + ("foo bar" + ""))) + "foo bar")))) {
                    case "\n\t":
                        if true {
                            if (4294967295 == 713) {
                                fmt.Println(fi15)
                                if true {
                                    var v26 int8 = -76
                                    if true {
                                        if true {
                                            var v27 float32 = 0.0
                                            var v28 int32 = 818
                                            var v29 uint8 = uint8(v1)
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 int16 = 548
                                            _ = v30
                                        }
                                        var v31 bool = true
                                        _ = v31
                                    }
                                    fmt.Println(v9)
                                    _ = v26
                                } else {
                                    var v32 bool = true
                                    _ = v32
                                }
                                _ = len(v5)
                            }
                        }
                    case "hello":
                        v4 += v4
                    case "":
                        v1 -= v1
                    }
                    switch (("hello" + "") + (("\n\t" + (("a" + ((("" + "a") + (("hello" + "foo bar") + "foo bar")) + ("\n\t" + "hello"))) + "foo bar")) + "a")) {
                    case "":
                        v4 -= (^(v4))
                    case "foo bar":
                        fmt.Println(v10)
                    case "\n\t":
                        v33 := GenS1{F0: 553, F1: 1}
                        _ = v33
                    }
                }
                v34 := GenS1{F0: v1, F1: 65535}
                _ = v34
            } else {
                if true {
                    if (65535 == 0) {
                        v1 = v1
                        var v35 int64 = genFn0()
                        _ = v35
                    }
                    v36 := GenS2{F0: 83}
                    _ = v36
                } else {
                    switch ("a" + "a") {
                    case "hello":
                        fmt.Println(v1, v10, fi11)
                    default:
                        v10 = (fi15 ^ (v10 << ((42) & 63)))
                        if (!true) {
                            v10 -= (v10 >> ((344) & 63))
                            var v37 uint32 = 4294967295
                            _ = v37
                        } else {
                            v38 := GenS2{F0: 100}
                            _ = v38
                        }
                    }
                    fmt.Println(fi11, v10, fi11)
                    delete(v5, 0)
                }
                if true {
                    _ = len(v5)
                } else {
                    v39 := GenS2{F0: 0}
                    var v40 int16 = 32767
                    _ = v39
                    _ = v40
                }
            }
            _ = fi15
        }
        _ = fi11
        _ = v14
    }
    v41 := GenS3{F0: (-(v10)), F1: (v4 >> ((1) & 63)), F2: (v1 % ((v1) | 1))}
    v4 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v9
}
