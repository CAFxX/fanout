package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 string; F1 int64; F2 int64 }
type GenS3 struct { F0 int; F1 int8; F2 uint }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint; F2 int }

func main() {
    v1 := make(map[string]string)
    var v2 int = 32768
    var v3 uint = 401
    var v4 uint64 = 2
    v5 := make(map[string]string)
    v4 = 91
    var v6 float32 = (-2.25 * 1.5)
    if (true || true) {
        v7 := GenS1{F0: 32767, F1: 0}
        _ = v7
    }
    v8 := GenS1{F0: 283, F1: uint8(v3)}
    v9 := v5["\n\t"]
    v4 += (v4 >> ((1) & 63))
    var v10 uint16 = 65535
    var v11 uint = v3
    for fi12 := 0; fi12 < 11; fi12++ {
        if false {
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 int64 = -133
                if true {
                    v15 := GenS1{F0: 42, F1: 42}
                    delete(v1, "foo bar")
                    if (!false) {
                        v16 := GenS2{F0: v9, F1: 9223372036854775807, F2: 210}
                        _ = v16
                    } else {
                        if ((("hello" + "a") + ("a" + (v9 + "\n\t"))) > (v9 + v9)) {
                            if (!(!(1.5 > (float64(v6) * float64(v6))))) {
                                v3 = (^(v11))
                                var v17 int16 = 42
                                _ = v17
                            }
                            v9 = ("\n\t" + "\n\t")
                            var v18 int = (fi13 | fi13)
                            _ = v18
                        }
                        v19 := []int{}
                        fmt.Println(fi12)
                        _ = v19
                    }
                    _ = v15
                } else {
                    if ((false || (!false)) || (false && ((v4 % ((v4) | 1)) < (v4 + v4)))) {
                        var v20 int16 = 42
                        v21 := GenS3{F0: 1, F1: 127, F2: (v11 / ((uint(v10)) | 1))}
                        _ = v20
                        _ = v21
                    } else {
                        v22 := GenS1{F0: -84, F1: 97}
                        var v23 int64 = -1
                        var v24 uint64 = uint64(v3)
                        if (!(v9 < v9)) {
                            v25 := []int{fi13}
                            if (1 <= 2) {
                                fmt.Println(v14, v10, v2)
                                v26 := GenS4{F0: (v11 + v11), F1: 128}
                                var v27 uint = (v3 / (((v3 | v3)) | 1))
                                var v28 uint64 = v24
                                v29 := GenS4{F0: (v11 | v3), F1: 12}
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                var v30 int = (fi12 % ((int(v24)) | 1))
                                _ = v30
                            }
                            _ = v25
                        } else {
                            v31 := GenS5{F0: (v10 | 65535), F1: (fi12 + (v2 >> ((v11) & 63)))}
                            _ = v31
                        }
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    v32 := v5[v9]
                    fmt.Println(v2, fi12, v10)
                    _ = v32
                }
                var v33 uint16 = v10
                _ = fi13
                _ = v14
                _ = v33
            }
            for fi34 := 0; fi34 < 2; fi34++ {
                if (((!false) || true) || (v11 < 0)) {
                    v4 = (v4 % (((v4 * (v4 >> (((v3 % ((v11) | 1))) & 63)))) | 1))
                    if false {
                        var v35 int8 = 127
                        var v36 string = v9
                        v36 = (v9 + v9)
                        _ = v35
                        _ = v36
                    } else {
                        fmt.Println(fi34, v10, fi12)
                        v37 := GenS6{F0: "\n\t", F1: 521}
                        if (42 == 829) {
                            v38 := GenS7{F0: 32767}
                            v39 := GenS1{F0: -32768, F1: 0}
                            var v40 uint64 = (v4 | 1)
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        _ = v37
                    }
                } else {
                    v41 := []int{}
                    var v42 float64 = (1e10 / 0.1)
                    v43 := GenS6{F0: ("" + "a"), F1: 1}
                    v44 := GenS5{F0: (v10 ^ 760), F1: -664}
                    v42 = v42
                    v3 -= v3
                    _ = v41
                    _ = v42
                    _ = v43
                    _ = v44
                }
                var v45 float32 = (3.14159 + 3.14159)
                _ = fi34
                _ = v45
            }
            fmt.Println(fi12)
        }
        if false {
            v11 -= (^(v11))
        } else {
            if true {
                if false {
                    if true {
                        v11 += (^(v11))
                    } else {
                        v46 := GenS8{F0: -1, F1: 805, F2: (v2 | v2)}
                        v1[("hello" + ((v9 + v9) + (v9 + "a")))] = "a"
                        _ = v46
                    }
                    var v47 int32 = -1
                    _ = v47
                }
            }
            v48 := GenS5{F0: (v10 << ((65535) & 63)), F1: 9223372036854775807}
            _ = v48
        }
        v3 += v11
        _ = fi12
    }
    v49 := []int{127, v2, (v2 << (((v3 - (v3 + 962))) & 63)), (v2 * v2)}
    v50 := GenS5{F0: (v10 - (v10 * (v10 >> (((v3 << (((v3 & v11)) & 63))) & 63)))), F1: 196}
    v3 -= v3
    for fi51 := 0; fi51 < 12; fi51++ {
        fmt.Println(v3, v9)
        var v52 uint64 = (v4 | 0)
        delete(v5, "foo bar")
        _ = fi51
        _ = v52
    }
    v4 += 2
    if (!true) {
        v3 = v3
        fmt.Println(v6, v9)
    } else {
        var v53 int16 = 231
        _ = v53
    }
    v54 := GenS8{F0: -737, F1: (v11 & uint(v10)), F2: 128}
    fmt.Println(v3)
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v54
    _ = v6
    _ = v8
    _ = v9
}
