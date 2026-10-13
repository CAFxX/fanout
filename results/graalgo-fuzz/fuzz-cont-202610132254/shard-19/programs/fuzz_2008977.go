package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint; F2 int32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8, p1 int64, p2 int64) bool {
    v6 := []int{42, 39}
    var v7 int16 = 338
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (false && (v7 != (v7 % ((32767) | 1))))
}
func genFn1() int32 {
    var v8 uint16 = 634
    var v9 uint8 = 1
    if (-36 < 100) {
        for fi10 := 0; fi10 < 6; fi10++ {
            if ((v8 | 0) < 42) {
                var v11 uint = 1
                for fi12 := 0; fi12 < 19; fi12++ {
                    v11 += (v11 ^ v11)
                    v13 := []int{127, (fi12 >> ((42) & 63)), -2147483648, (fi12 % ((fi10) | 1))}
                    _ = fi12
                    _ = v13
                }
                switch (fi10 % ((42) | 1)) {
                case 2147483647:
                    var v14 uint64 = 4294967295
                    _ = v14
                }
                _ = v11
            } else {
                v9 = v9
                v8 += 42
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS1{F0: "foo bar"}
                    if false {
                        var v17 uint16 = (v8 % (((v8 >> ((0) & 63))) | 1))
                        var v18 uint32 = 2147483648
                        _ = v17
                        _ = v18
                    } else {
                        var v19 bool = (!(false || (32767 < 32767)))
                        switch (fi10 + (fi10 << ((65535) & 63))) {
                        case -1:
                            v8 = 495
                            v8 -= 65535
                            v20 := []int{32767, fi15, 42, fi15}
                            var v21 float32 = (1e10 / 3.14159)
                            _ = v20
                            _ = v21
                        case 0:
                            var v22 float32 = (0.0 + (0.1 * 0.0))
                            _ = v22
                        case -9223372036854775808:
                            var v23 int16 = 32767
                            v8 = (v8 << ((18446744073709551615) & 63))
                            v24 := GenS2{F0: -1}
                            _ = v23
                            _ = v24
                        }
                        fmt.Println(v19, v19, v19)
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                }
            }
            switch "a" {
            case "a":
                fmt.Println(v8, v9, v9)
            case "foo bar":
                fmt.Println(v8, v8)
            case "\n\t":
                v25 := []int{fi10, fi10}
                _ = v25
            default:
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 uint = 471
                    switch "hello" {
                    case "foo bar":
                        v28 := GenS2{F0: 17}
                        _ = v28
                    case "":
                        var v29 uint64 = 259
                        _ = v29
                    case "hello":
                        fmt.Println(fi10)
                    default:
                        if (65535 < 663) {
                            var v30 int16 = -342
                            _ = v30
                        } else {
                            v31 := []int{32767}
                            var v32 uint8 = (v9 | (v9 ^ (v9 - (^(v9)))))
                            _ = len(v31)
                            _ = v31
                            _ = v32
                        }
                    }
                    v8 += (v8 << (((v27 ^ v27)) & 63))
                    _ = fi26
                    _ = v27
                }
            }
            _ = fi10
        }
        if true {
            for fi33 := 0; fi33 < 8; fi33++ {
                v9 -= (v9 | (^(v9)))
                var v34 string = "a"
                var v35 uint64 = 2
                _ = fi33
                _ = v34
                _ = v35
            }
            var v36 uint8 = v9
            if (32767 == 0) {
                v36 = v9
                var v37 int32 = 1
                v38 := GenS3{F0: 42, F1: 410, F2: (v37 | (v37 / ((0) | 1)))}
                _ = v37
                _ = v38
            } else {
                fmt.Println(v9, v36, v36)
                v36 += (v36 << ((0) & 63))
            }
            _ = v36
        }
        v39 := GenS4{F0: 1}
        _ = v39
    }
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 int8 = 100
    var v5 float64 = float64((1e10 + (1e10 - float32(1.5))))
    v1 -= v1
    var v40 uint64 = 823
    v40 = 627
    for fi41 := 0; fi41 < 20; fi41++ {
        for fi42 := 0; fi42 < 6; fi42++ {
            v5 += (0.0 * (v1 / (float64((0.1 - (0.1 * 1.5))) / 1.0)))
            _ = fi42
        }
        _ = fi41
    }
    var v43 int64 = 42
    fmt.Println(v40)
    v4 = -1
    var v44 uint64 = v40
    switch "" {
    case "":
        for fi45 := 0; fi45 < 8; fi45++ {
            if true {
                v46 := []int{fi45}
                _ = v46
            }
            for fi47 := 0; fi47 < 9; fi47++ {
                switch fi45 {
                case 9223372036854775807:
                    fmt.Println(v44, v5, v43)
                }
                _ = fi47
            }
            _ = fi45
        }
    case "hello":
        fmt.Println(v44, v40, v44)
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 uint32 = 2147483648
        var v50 string = (("\n\t" + "\n\t") + "a")
        _ = fi48
        _ = v49
        _ = v50
    }
    var v51 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v51
}
