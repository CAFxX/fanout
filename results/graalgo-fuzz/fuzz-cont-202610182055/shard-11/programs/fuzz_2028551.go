package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int64 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint64, p2 int32) int16 {
    fmt.Println(p0, p0, p2)
    fmt.Println(p0, p1, p2)
    var v7 int = 83
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 int16 = 42
    var v2 uint = 224
    v3 := make(map[string]uint32)
    var v4 int32 = -87
    v5 := make(map[uint32]int)
    var v6 int32 = (^(v4))
    var v8 uint8 = 127
    v2 = (v2 % ((370) | 1))
    var v9 uint16 = 32767
    var v10 uint16 = v9
    var v11 uint64 = 906
    for fi12 := 0; fi12 < 17; fi12++ {
        v2 += uint(v8)
        _ = fi12
    }
    fmt.Println(v9, v2, v9)
    fmt.Println(v4, v2)
    for fi13 := 0; fi13 < 15; fi13++ {
        switch (-(fi13)) {
        case 4294967295:
            var v14 uint32 = 65535
            _ = v14
        case 128:
            fmt.Println(v2, v10, v11)
        case 32767:
            v5[1] = fi13
        default:
            v15 := []int{}
            _ = v15
        }
        _ = fi13
    }
    var v16 int64 = 42
    if false {
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 float32 = 1e10
            if ((!(true || false)) && true) {
                for fi19 := 0; fi19 < 15; fi19++ {
                    fmt.Println(v2)
                    v9 = v10
                    var v20 int16 = (v1 ^ v1)
                    _ = fi19
                    _ = v20
                }
                v21 := v3["\n\t"]
                v22 := GenS1{F0: (v8 ^ 2), F1: 42, F2: v16}
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                v24 := GenS2{F0: true, F1: fi17}
                _ = fi23
                _ = v24
            }
            _ = fi17
            _ = v18
        }
        if (623 == 42) {
            switch (("foo bar" + "hello") + ("hello" + "")) {
            case "foo bar":
                var v25 int8 = 100
                _ = v25
            default:
                if false {
                    v8 += 160
                }
            }
        }
        fmt.Println(v10, v9, v9)
    }
    if true {
        if ((v16 << ((4294967295) & 63)) >= -9223372036854775808) {
            if false {
                if (v16 < v16) {
                    switch 2147483647 {
                    case 0:
                        v6 += (v6 >> ((v2) & 63))
                    case 32768:
                        if true {
                            if (!(!true)) {
                                var v26 uint = v2
                                _ = v26
                            } else {
                                v10 -= v10
                                for fi27 := 0; fi27 < 20; fi27++ {
                                    var v28 bool = (2147483647 <= 837)
                                    v28 = true
                                    _ = fi27
                                    _ = v28
                                }
                                v29 := GenS3{F0: v9}
                                _ = v29
                            }
                            var v30 int8 = 25
                            var v31 int8 = v30
                            _ = v30
                            _ = v31
                        }
                    case 2147483647:
                        if (int(v1) > 32767) {
                            for fi32 := 0; fi32 < 1; fi32++ {
                                v33 := GenS3{F0: 255}
                                var v34 int32 = v6
                                var v35 int32 = -67
                                _ = fi32
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            var v36 int64 = (-(v16))
                            _ = v36
                        }
                    default:
                        for fi37 := 0; fi37 < 3; fi37++ {
                            v16 -= (v16 >> ((65535) & 63))
                            _ = fi37
                        }
                    }
                } else {
                    v4 = v6
                    for fi38 := 0; fi38 < 1; fi38++ {
                        switch "" {
                        case "":
                            switch fi38 {
                            case 4294967295:
                                for fi39 := 0; fi39 < 14; fi39++ {
                                    v2 += v2
                                    _ = fi39
                                }
                            case 32767:
                                if true {
                                    v40 := GenS2{F0: ((^(v8)) <= (v8 * 255)), F1: fi38}
                                    _ = v40
                                }
                            case -2147483648:
                                v2 = uint(v11)
                            }
                        case "a":
                            for fi41 := 0; fi41 < 15; fi41++ {
                                v1 = v1
                                v3[""] = uint32(v2)
                                fmt.Println(v10, v8)
                                _ = fi41
                            }
                        default:
                            for fi42 := 0; fi42 < 11; fi42++ {
                                var v43 uint = 0
                                _ = fi42
                                _ = v43
                            }
                        }
                        _ = fi38
                    }
                    v44 := GenS3{F0: 909}
                    _ = v44
                }
                v4 = v4
                v9 -= (v9 / (((v10 + v10)) | 1))
            }
            v3["hello"] = 791
            v45 := v5[758]
            _ = v45
        } else {
            fmt.Println(v10)
            var v46 int64 = v16
            _ = v46
        }
        var v47 int16 = -622
        if (!((true || (!true)) && true)) {
            v48 := v3["foo bar"]
            fmt.Println(v47)
            _ = v48
        } else {
            v3["a"] = 0
        }
        _ = v47
    }
    v2 -= v2
    var v49 int32 = v4
    var v50 uint16 = (v9 | 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v8
    _ = v9
}
