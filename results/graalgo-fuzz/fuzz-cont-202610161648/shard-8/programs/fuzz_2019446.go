package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
type GenS3 struct { F0 string; F1 int; F2 int }
func genFn0(p0 int64, p1 uint16, p2 uint64) bool {
    p2 -= 2
    p1 = (p1 >> ((18446744073709551615) & 63))
    var v7 uint = 1
    v8 := GenS1{F0: 348, F1: 32767, F2: 65535}
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -1
    var v3 int = 2147483647
    var v4 string = "\n\t"
    v5 := make(map[int]string)
    v6 := make(map[uint32]int)
    fmt.Println(v2, v4)
    v9 := GenS1{F0: 325, F1: 521, F2: 395}
    if false {
        if true {
            v2 = 0
            var v10 int64 = 9223372036854775807
            _ = v10
        } else {
            v4 = "a"
            switch v4 {
            case "foo bar":
                fmt.Println(v4, v4)
            default:
                for fi11 := 0; fi11 < 19; fi11++ {
                    var v12 string = v4
                    _ = fi11
                    _ = v12
                }
            }
            v2 -= 127
        }
        fmt.Println(v3, v4, v3)
    }
    var v13 int64 = 0
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint16 = 113
        _ = fi14
        _ = v15
    }
    v16 := v5[(v3 + (v3 + (v3 | v3)))]
    v17 := GenS1{F0: v13, F1: -1, F2: 2}
    switch v3 {
    case 9223372036854775807:
        for fi18 := 0; fi18 < 16; fi18++ {
            if ((v13 >> ((255) & 63)) > v13) {
                if (782 <= 1) {
                    v19 := []int{fi18, (v3 | fi18), (fi18 * (v3 | (fi18 >> ((85) & 63))))}
                    _ = v19
                }
            }
            if (true || (3.14159 <= 0.0)) {
                fmt.Println(v13)
                v3 += v3
                v20 := GenS1{F0: v13, F1: 529, F2: 2}
                _ = v20
            }
            if false {
                for fi21 := 0; fi21 < 11; fi21++ {
                    fmt.Println(v4, v4, v4)
                    _ = fi21
                }
                var v22 int64 = (v13 % ((815) | 1))
                switch (v16 + v16) {
                case "":
                    var v23 uint8 = 2
                    _ = v23
                case "a":
                    v24 := v1[32]
                    _ = v24
                default:
                    for fi25 := 0; fi25 < 16; fi25++ {
                        switch fi25 {
                        case 128:
                            v26 := []int{fi18, fi25, (fi18 * fi25), v3}
                            _ = v26
                        default:
                            v13 -= v13
                        }
                        _ = fi25
                    }
                }
                _ = v22
            }
            _ = fi18
        }
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        if (128 != 2) {
            for fi28 := 0; fi28 < 5; fi28++ {
                if false {
                    var v29 int64 = 0
                    var v30 uint = 2
                    _ = v29
                    _ = v30
                } else {
                    var v31 int8 = v2
                    v32 := GenS2{F0: v2, F1: (fi27 >> ((18446744073709551615) & 63)), F2: 255}
                    v33 := GenS1{F0: v13, F1: -413, F2: 2}
                    var v34 string = ((v4 + (v16 + "")) + (("hello" + "a") + v16))
                    _ = v31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                if (255 < 255) {
                    delete(v1, 1)
                } else {
                    var v35 uint = 65535
                    fmt.Println(v16, fi28, v2)
                    var v36 uint16 = 568
                    _ = v35
                    _ = v36
                }
                _ = fi28
            }
        }
        fmt.Println(fi27)
        v37 := GenS2{F0: v2, F1: (v3 | v3), F2: 1}
        _ = fi27
        _ = v37
    }
    v38 := GenS3{F0: v16, F1: (v3 << ((343) & 63)), F2: v3}
    v4 = v4
    v39 := GenS3{F0: ("a" + v16), F1: int(v2), F2: (v3 | (-(v3)))}
    v40 := GenS1{F0: v13, F1: 42, F2: 2}
    var v41 int64 = (-(v13))
    var v42 int64 = (-(v41))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v9
}
