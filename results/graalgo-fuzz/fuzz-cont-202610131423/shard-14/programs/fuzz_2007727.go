package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32 }
func genFn0() uint8 {
    var v5 uint64 = 626
    if ((0 != 65535) && (true && (false && (2 > 0)))) {
        v6 := []int{-1}
        _ = v6
    } else {
        if true {
            v7 := GenS1{F0: 9223372036854775807}
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            v5 = v5
            fmt.Println(v5)
            switch "\n\t" {
            case "a":
                v10 := GenS1{F0: 1}
                _ = v10
            case "foo bar":
                v11 := GenS2{F0: 343}
                _ = v11
            case "":
                if false {
                    switch "" {
                    case "hello":
                        var v12 bool = true
                        _ = v12
                    default:
                        var v13 uint16 = 1
                        _ = v13
                    }
                } else {
                    var v14 bool = true
                    v15 := GenS3{F0: 42, F1: 42}
                    v5 = 42
                    _ = v14
                    _ = v15
                }
            }
            _ = fi9
        }
        v16 := []int{fi8}
        _ = fi8
        _ = v16
    }
    v5 += 0
    if true {
        v17 := []int{127, 128, -23}
        v18 := GenS1{F0: -9223372036854775808}
        _ = v17
        _ = v18
    } else {
        v19 := []int{1, -9223372036854775808, -2147483648}
        if (false || (-32768 != 32767)) {
            v5 -= 1
            v20 := GenS3{F0: 1, F1: 4294967295}
            _ = v20
        }
        fmt.Println(v5)
        _ = v19
    }
    _ = v5
    return 187
}

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 1
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    v2 = v2
    var v21 int32 = 42
    v22 := GenS1{F0: 0}
    fmt.Println(v1, v21, v1)
    for fi23 := 0; fi23 < 20; fi23++ {
        switch -853 {
        case 127:
            v1 -= (v1 - uint8(v2))
        case 1:
            if ((((3.14159 + float32(1.5)) / -2.25) / (-2.25 / 1.0)) > ((float32(float64(0.1)) - -2.25) - 0.1)) {
                v1 += (v1 + v1)
            } else {
                v1 = v1
                for fi24 := 0; fi24 < 9; fi24++ {
                    if (((702 <= uint(v1)) || (242 > v1)) && false) {
                        fmt.Println(v2)
                        fmt.Println(fi24)
                        v2 += 65535
                        fmt.Println(v1)
                        if false {
                            _ = len(v4)
                            var v25 float32 = 1e10
                            v26 := GenS2{F0: 18446744073709551615}
                            _ = v25
                            _ = v26
                        }
                        if false {
                            switch (((("\n\t" + "foo bar") + "a") + "") + ("\n\t" + "foo bar")) {
                            case "":
                                v21 -= v21
                            }
                            v2 -= v2
                        }
                        v3[v2] = "foo bar"
                        if (true || false) {
                            v21 = v21
                            v21 -= v21
                            v21 = 2147483647
                        } else {
                            v27 := GenS1{F0: 42}
                            var v28 int32 = (v21 * (v21 << ((42) & 63)))
                            fmt.Println(v1, v28)
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS2{F0: 2}
                        v30 := GenS3{F0: v1, F1: 255}
                        v1 += (v1 ^ (v1 / ((209) | 1)))
                        fmt.Println(v21)
                        _ = v29
                        _ = v30
                    } else {
                        var v31 bool = (false || true)
                        var v32 int32 = v21
                        var v33 uint64 = 46
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v34 := GenS3{F0: v1, F1: 484}
                    v1 = (^(v1))
                    _ = fi24
                    _ = v34
                }
            }
        }
        v1 = genFn0()
        _ = fi23
    }
    v35 := GenS4{F0: -9223372036854775808, F1: "foo bar"}
    var v36 int64 = -392
    if false {
        v37 := []int{620, 32767, -1, 42}
        var v38 float32 = 1.5
        var v39 uint16 = 0
        _ = v37
        _ = v38
        _ = v39
    }
    var v40 uint16 = 255
    v40 = (v40 << ((910) & 63))
    for fi41 := 0; fi41 < 3; fi41++ {
        for fi42 := 0; fi42 < 1; fi42++ {
            if ((!false) || true) {
                v43 := GenS3{F0: (v1 * v1), F1: 255}
                v44 := GenS3{F0: 128, F1: 0}
                v1 -= v1
                v45 := GenS5{F0: 0, F1: 42}
                _ = v43
                _ = v44
                _ = v45
            }
            v46 := GenS6{F0: 2147483648}
            v2 = (v2 / (((v2 + v2)) | 1))
            _ = fi42
            _ = v46
        }
        for fi47 := 0; fi47 < 5; fi47++ {
            v48 := GenS7{F0: uint32(v1)}
            if false {
                v49 := []int{(fi41 - (fi41 * -9223372036854775808)), 2147483647, (fi47 * -2147483648), fi41}
                if ("hello" != "") {
                    delete(v3, v2)
                }
                v50 := GenS6{F0: uint32(v40)}
                _ = v49
                _ = v50
            } else {
                var v51 float32 = (0.1 / 0.1)
                var v52 int32 = v21
                if (false && true) {
                    v4[(^(v2))] = v2
                    if false {
                        if true {
                            fmt.Println(v51, fi41)
                        }
                        var v53 uint64 = 1
                        v21 = (v52 >> ((uint(v1)) & 63))
                        _ = v53
                    }
                }
                _ = v51
                _ = v52
            }
            v54 := GenS6{F0: (v2 ^ (^(v2)))}
            _ = fi47
            _ = v48
            _ = v54
        }
        v55 := []int{fi41, fi41, -1}
        _ = fi41
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
}
