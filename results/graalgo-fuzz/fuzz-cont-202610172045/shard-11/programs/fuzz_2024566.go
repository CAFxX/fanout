package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int }
func genFn0(p0 uint, p1 uint32) int8 {
    v7 := []int{-1, 764}
    if false {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        for fi8 := 0; fi8 < 7; fi8++ {
            if (0 != 232) {
                var v9 int64 = -9223372036854775808
                _ = v9
            }
            var v10 uint32 = (p1 & (p1 + p1))
            _ = fi8
            _ = v10
        }
    } else {
        p1 = (p1 << (((p0 * p0)) & 63))
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 uint8, p2 uint64) int16 {
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: fi11, F1: 960}
        if true {
            var v13 uint8 = p1
            v14 := GenS2{F0: 42, F1: 2}
            if true {
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := GenS2{F0: uint64(p1), F1: 65535}
                    _ = fi15
                    _ = v16
                }
                switch fi11 {
                case 1:
                    var v17 string = "a"
                    _ = v17
                }
                if (!true) {
                    v18 := []int{fi11, 32767, (fi11 % ((fi11) | 1))}
                    p2 -= (p2 - p2)
                    switch "" {
                    case "":
                        fmt.Println(p2)
                    }
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 5; fi19++ {
                    var v20 string = ""
                    fmt.Println(v20, p0, p0)
                    _ = fi19
                    _ = v20
                }
                var v21 int = 127
                var v22 uint32 = 1
                _ = v21
                _ = v22
            }
            _ = v13
            _ = v14
        } else {
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := []int{fi11, fi23}
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 14; fi25++ {
                var v26 float64 = 1e10
                _ = fi25
                _ = v26
            }
        }
        _ = fi11
        _ = v12
    }
    var v27 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return -32768
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 57
    var v3 uint16 = 32767
    var v4 int = -1
    var v5 uint8 = 127
    var v6 uint16 = (v3 | 0)
    v4 = 32768
    if (1e10 == (0.0 + -2.25)) {
        v28 := []int{-813, (-(v4))}
        _ = v28
    }
    delete(v1, v2)
    var v29 uint16 = v6
    v30 := GenS1{F0: 881, F1: (v2 & v2)}
    for fi31 := 0; fi31 < 11; fi31++ {
        v32 := GenS1{F0: (fi31 * v4), F1: (v2 >> ((42) & 63))}
        _ = fi31
        _ = v32
    }
    if true {
        if (-601 < -395) {
            switch (v4 * -9223372036854775808) {
            case 1:
                v29 = (v3 * v6)
            case 32767:
                if false {
                    if (1000 == 1000) {
                        v1[(^(v2))] = true
                    }
                    delete(v1, (v2 & v2))
                }
            default:
                v33 := GenS2{F0: 42, F1: (v29 / ((v6) | 1))}
                _ = v33
            }
            switch "foo bar" {
            case "hello":
                var v34 uint32 = v2
                _ = v34
            case "foo bar":
                delete(v1, (v2 >> ((42) & 63)))
            case "a":
                v35 := GenS2{F0: 488, F1: v3}
                _ = v35
            }
        }
        v36 := GenS3{F0: (float64(0.0) > (0.0 * 0.1)), F1: 0}
        _ = v36
    } else {
        if false {
            v37 := GenS4{F0: (v4 % ((int(v29)) | 1))}
            if (!(((8 <= 42) && (1 <= genFn1(1, v5, 2))) || (true && (!(0.0 >= (0.1 - 0.1)))))) {
                for fi38 := 0; fi38 < 1; fi38++ {
                    var v39 uint32 = 4294967295
                    _ = fi38
                    _ = v39
                }
            } else {
                v29 = 256
            }
            _ = v37
        } else {
            switch (v4 | (-(v4))) {
            case -9223372036854775808:
                var v40 uint32 = (v2 % ((686) | 1))
                _ = v40
            case -860:
                fmt.Println(v29)
            case 132:
                v3 -= v29
            default:
                var v41 float64 = 1.5
                _ = v41
            }
        }
        for fi42 := 0; fi42 < 15; fi42++ {
            for fi43 := 0; fi43 < 5; fi43++ {
                _ = len(v1)
                if (!false) {
                    var v44 uint8 = 127
                    _ = v44
                }
                _ = fi43
            }
            _ = fi42
        }
    }
    var v45 int64 = -909
    v46 := GenS2{F0: 2, F1: v29}
    var v47 float32 = 1.5
    v48 := GenS4{F0: v4}
    v49 := GenS3{F0: (127 < 8), F1: 201}
    for fi50 := 0; fi50 < 14; fi50++ {
        switch "\n\t" {
        case "\n\t":
            v51 := GenS1{F0: 32768, F1: 768}
            _ = v51
        }
        _ = fi50
    }
    fmt.Println(v2)
    v52 := []int{(v4 >> ((0) & 63)), (v4 | v4), v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v6
}
