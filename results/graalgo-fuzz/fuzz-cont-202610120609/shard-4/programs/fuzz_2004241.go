package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint16, p1 string) bool {
    if false {
        switch "foo bar" {
        case "":
            p0 = (p0 + (p0 | (^(p0))))
        default:
            v4 := []int{127}
            _ = v4
        }
        v5 := GenS1{F0: -1, F1: 127}
        _ = v5
    } else {
        fmt.Println(p0, p1)
        switch 32767 {
        case 2147483647:
            var v6 uint64 = 566
            _ = v6
        case -351:
            v7 := GenS1{F0: 42, F1: 765}
            _ = v7
        case 32768:
            v8 := []int{453, 127, 127, 9223372036854775807}
            _ = v8
        }
        var v9 int8 = 100
        _ = v9
    }
    v10 := GenS2{F0: 1}
    var v11 uint64 = 4294967295
    var v12 bool = ((p1 > ("" + p1)) && ((-2.25 < ((-2.25 / float32((float64(1.5) + 3.14159))) / 1.0)) && true))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = 9223372036854775807
    var v3 int8 = 127
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint32 = v1
        fmt.Println(v1, v2)
        switch "" {
        case "foo bar":
            v14 = (v14 << ((4294967295) & 63))
        case "":
            var v15 uint64 = 486
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int16 = 32767
        var v18 int64 = v2
        v19 := []int{}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 string = ((("foo bar" + "") + ("hello" + ("hello" + ("a" + ("hello" + "hello"))))) + "hello")
    v3 = (v3 + 81)
    for fi21 := 0; fi21 < 12; fi21++ {
        switch fi21 {
        case 1:
            v22 := []int{}
            _ = v22
        case -1:
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 float32 = (1.5 * float32((-2.25 - (3.14159 - (0.1 + 0.0)))))
                v25 := []int{fi23, fi23, (fi23 / (((fi21 & (fi23 | (fi23 % ((32767) | 1))))) | 1)), (fi23 % (((fi23 | fi21)) | 1))}
                switch "a" {
                case "foo bar":
                    var v26 int32 = 2147483647
                    _ = v26
                case "a":
                    v27 := GenS1{F0: 1000, F1: fi23}
                    _ = v27
                case "\n\t":
                    fmt.Println(v2, v24)
                }
                _ = fi23
                _ = v24
                _ = v25
            }
        case 401:
            switch fi21 {
            case 42:
                if false {
                    var v28 int = 127
                    v29 := GenS1{F0: 535, F1: 127}
                    _ = v28
                    _ = v29
                }
            case 32768:
                var v30 uint64 = 642
                _ = v30
            }
        default:
            switch fi21 {
            case 779:
                switch (^(fi21)) {
                case 42:
                    v31 := GenS3{F0: 678}
                    _ = v31
                case -800:
                    var v32 uint32 = v1
                    _ = v32
                default:
                    v33 := []int{}
                    _ = v33
                }
            case -405:
                v34 := GenS2{F0: v2}
                _ = v34
            case 4294967295:
                var v35 int32 = 0
                _ = v35
            }
        }
        _ = fi21
    }
    if (32767 == 32767) {
        v36 := GenS4{F0: -1}
        _ = v36
    }
    if (934 == 776) {
        for fi37 := 0; fi37 < 17; fi37++ {
            var v38 uint = 1
            _ = fi37
            _ = v38
        }
        var v39 uint = 1
        _ = v39
    }
    v40 := GenS5{F0: 34}
    v41 := GenS1{F0: -36, F1: 42}
    v3 = (v3 | (v3 - v3))
    v42 := []int{643, 32768, 949, 127}
    if true {
        var v43 uint8 = 127
        for fi44 := 0; fi44 < 20; fi44++ {
            if len(v42) > 0 {
                v42[0] = fi44
            }
            _ = fi44
        }
        var v45 uint16 = 32767
        _ = v43
        _ = v45
    } else {
        for fi46 := 0; fi46 < 15; fi46++ {
            for fi47 := 0; fi47 < 4; fi47++ {
                var v48 uint = 4294967295
                fmt.Println(v20, v3)
                v2 -= v2
                _ = fi47
                _ = v48
            }
            _ = fi46
        }
        if (42 != 871) {
            for fi49 := 0; fi49 < 17; fi49++ {
                if (0 == 840) {
                    v50 := GenS3{F0: uint64(v1)}
                    var v51 uint16 = 891
                    if len(v42) > 0 {
                        v42[0] = fi49
                    }
                    _ = v50
                    _ = v51
                } else {
                    fmt.Println(fi49)
                    v2 = (v2 / ((1) | 1))
                }
                _ = fi49
            }
            if len(v42) >= 2 {
                v42 = v42[1:len(v42)]
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v40
    _ = v41
    _ = v42
}
