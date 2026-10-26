package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int; F2 int16 }
type GenS4 struct { F0 string; F1 bool; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint8, p1 string) int16 {
    p1 = p1
    var v7 uint32 = 65535
    var v8 bool = (!false)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int8 = 42
    var v2 int = 42
    var v3 float64 = (0.0 - float64(0.0))
    var v4 uint16 = 65535
    v5 := make(map[string]bool)
    v6 := make(map[int]string)
    var v9 float64 = v3
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: 42, F1: 2147483648, F2: 65535}
            var v13 uint64 = 8
            if true {
                v1 = (v1 * v1)
                fmt.Println(v1, v9, v2)
                v14 := GenS1{F0: 2, F1: 319, F2: 756}
                _ = v14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS2{F0: (((v1 ^ (^(v1))) != 127) && (454 != 42))}
        _ = fi15
        _ = v16
    }
    if false {
        v17 := []int{}
        var v18 uint16 = (v4 << ((65535) & 63))
        switch "" {
        case "foo bar":
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 int16 = 32767
                var v21 uint64 = 420
                var v22 uint32 = uint32(v4)
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        }
        _ = v17
        _ = v18
    } else {
        for fi23 := 0; fi23 < 11; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                fmt.Println(fi24)
                v25 := GenS1{F0: 128, F1: 1, F2: 1}
                var v26 string = ("foo bar" + "a")
                if (1 > 42) {
                    v9 -= ((v9 / (v9 / 1.0)) + (0.1 * v3))
                    var v27 int32 = -2147483648
                    fmt.Println(v27, v2, v2)
                    _ = v27
                }
                v28 := GenS3{F0: 0, F1: 128, F2: -311}
                _ = fi24
                _ = v25
                _ = v26
                _ = v28
            }
            var v29 bool = false
            _ = fi23
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        switch v2 {
        case 859:
            v31 := GenS4{F0: "", F1: (0 != -1), F2: v4}
            _ = v31
        case 218:
            v3 = v3
        default:
            v1 = (^(v1))
        }
        var v32 uint = 54
        v2 += fi30
        _ = fi30
        _ = v32
    }
    v1 += (v1 / (((v1 * v1)) | 1))
    switch v2 {
    case 185:
        switch (v2 * (v2 / ((4294967295) | 1))) {
        case -9223372036854775808:
            var v33 uint = 2
            _ = v33
        case 320:
            for fi34 := 0; fi34 < 2; fi34++ {
                v9 -= 0.0
                if true {
                    fmt.Println(v4)
                    v2 = -389
                    var v35 int64 = -9223372036854775808
                    _ = v35
                } else {
                    v1 = (^(v1))
                }
                var v36 int16 = 529
                _ = fi34
                _ = v36
            }
        default:
            v3 = v3
        }
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        switch ("\n\t" + "a") {
        case "a":
            v3 = 0.1
        case "\n\t":
            v38 := GenS5{F0: "foo bar"}
            _ = v38
        case "hello":
            var v39 int32 = 1
            _ = v39
        default:
            _ = len(v5)
        }
        fmt.Println(v3, v3, v9)
        v40 := GenS6{F0: 255, F1: v4}
        _ = fi37
        _ = v40
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        v42 := GenS3{F0: 1, F1: fi41, F2: 389}
        var v43 uint = 42
        v2 += -909
        _ = fi41
        _ = v42
        _ = v43
    }
    v9 = float64(0.1)
    for fi44 := 0; fi44 < 3; fi44++ {
        var v45 bool = false
        if false {
            v1 += v1
            fmt.Println(v2)
        } else {
            var v46 string = "foo bar"
            var v47 int32 = -963
            if (2147483648 >= 42) {
                v2 = v2
                fmt.Println(v4, v9)
            } else {
                switch v46 {
                case "a":
                    v46 = "\n\t"
                }
                for fi48 := 0; fi48 < 17; fi48++ {
                    v49 := []int{int(v47), v2, 1}
                    _ = fi48
                    _ = v49
                }
            }
            _ = v46
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    var v50 int16 = 1000
    v50 = genFn0(1, "")
    if true {
        if (false || true) {
            if (54 < (^(v1))) {
                v4 = v4
            }
        }
        var v51 uint = 65535
        _ = v51
    } else {
        for fi52 := 0; fi52 < 8; fi52++ {
            var v53 int32 = 680
            if true {
                v54 := []int{fi52, (v2 / ((42) | 1)), (^(v2))}
                fmt.Println(v53, fi52, fi52)
                var v55 int8 = (v1 + (-(v1)))
                _ = v54
                _ = v55
            }
            v56 := GenS7{F0: (v4 + 179)}
            _ = fi52
            _ = v53
            _ = v56
        }
    }
    var v57 uint16 = v4
    delete(v5, ("\n\t" + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v57
    _ = v6
    _ = v9
}
