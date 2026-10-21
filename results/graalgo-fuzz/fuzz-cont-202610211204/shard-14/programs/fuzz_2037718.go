package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int) int32 {
    if false {
        if (!(!true)) {
            if false {
                var v4 int16 = -1
                v5 := []int{(p0 - (p0 & (p0 * 32768))), p0, -140, 4294967295}
                _ = v4
                _ = v5
            } else {
                p0 = (p0 ^ (p0 * p0))
            }
            var v6 int64 = -886
            v7 := []int{(p0 | p0), (p0 - 499), -179}
            _ = v6
            _ = v7
        } else {
            var v8 int = (p0 + 2147483647)
            switch 32767 {
            case 128:
                v9 := []int{}
                _ = v9
            case 2147483647:
                var v10 float32 = 1.5
                _ = v10
            case 32767:
                fmt.Println(v8)
            }
            for fi11 := 0; fi11 < 2; fi11++ {
                switch "a" {
                case "":
                    var v12 int = v8
                    _ = v12
                }
                if (588 >= 4294967295) {
                    if true {
                        p0 = p0
                        if true {
                            var v13 bool = ((-114 <= 1) || (-9223372036854775808 <= -9223372036854775808))
                            _ = v13
                        }
                        fmt.Println(p0)
                    }
                }
                var v14 uint32 = 2147483648
                _ = fi11
                _ = v14
            }
            _ = v8
        }
        var v15 int = (p0 ^ (p0 - 32767))
        p0 += v15
        _ = v15
    } else {
        var v16 int = p0
        for fi17 := 0; fi17 < 13; fi17++ {
            v16 = 127
            switch p0 {
            case 4294967295:
                for fi18 := 0; fi18 < 2; fi18++ {
                    v16 = 4294967295
                    _ = fi18
                }
            }
            var v19 int16 = 0
            _ = fi17
            _ = v19
        }
        _ = v16
    }
    var v20 bool = (-1 >= -116)
    _ = p0
    _ = v20
    return 228
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    var v3 uint64 = 172
    v21 := GenS1{F0: 0, F1: "hello", F2: 2}
    switch 32768 {
    case 676:
        if (true && true) {
            var v22 uint = 255
            var v23 float32 = 3.14159
            var v24 int16 = -32768
            _ = v22
            _ = v23
            _ = v24
        }
    case 4294967295:
        var v25 int = -9223372036854775808
        _ = v25
    case 0:
        v1[910] = false
    default:
        if false {
            v3 -= 162
        } else {
            var v26 int64 = 697
            for fi27 := 0; fi27 < 8; fi27++ {
                v3 -= (v3 / (((v3 >> ((4294967295) & 63))) | 1))
                v28 := []int{int(v26), 127, (fi27 / ((fi27) | 1))}
                v26 = v26
                _ = fi27
                _ = v28
            }
            _ = v26
        }
    }
    v29 := []int{-864, 128, 32767, -1}
    _ = len(v29)
    fmt.Println(v3)
    v30 := GenS1{F0: 42, F1: ("a" + "foo bar"), F2: 42}
    if (!(true || false)) {
        if false {
            for fi31 := 0; fi31 < 20; fi31++ {
                for fi32 := 0; fi32 < 8; fi32++ {
                    fmt.Println(fi31, v3, fi31)
                    fmt.Println(fi32)
                    var v33 bool = (!false)
                    _ = fi32
                    _ = v33
                }
                fmt.Println(v3, fi31)
                v34 := GenS1{F0: 42, F1: "foo bar", F2: 676}
                _ = fi31
                _ = v34
            }
            if (3.14159 <= 1e10) {
                if (!(!false)) {
                    for fi35 := 0; fi35 < 11; fi35++ {
                        v36 := GenS1{F0: 32767, F1: "a", F2: 0}
                        _ = fi35
                        _ = v36
                    }
                    var v37 uint64 = (v3 << ((42) & 63))
                    _ = v37
                }
            } else {
                var v38 int32 = 1
                v39 := GenS1{F0: 0, F1: "a", F2: 2}
                _ = v38
                _ = v39
            }
            v40 := GenS2{F0: 925}
            _ = v40
        } else {
            v2[2147483647] = -9223372036854775808
        }
        fmt.Println(v3)
        var v41 int64 = 1
        _ = v41
    } else {
        v3 += (v3 >> ((2) & 63))
        v42 := v1[42]
        v43 := GenS3{F0: 852, F1: 0}
        _ = v42
        _ = v43
    }
    v3 += (v3 / ((1) | 1))
    var v44 uint8 = uint8(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v44
}
