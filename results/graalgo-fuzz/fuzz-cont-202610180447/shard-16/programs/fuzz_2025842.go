package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int; F2 string }
func genFn0(p0 int16) bool {
    p0 = (^(p0))
    var v4 string = ("\n\t" + "a")
    p0 += (p0 + (-(p0)))
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int64 = 0
        if ((false || true) && (!(!(1 > 1)))) {
            switch fi5 {
            case 1:
                v7 := []int{fi5, fi5, fi5}
                _ = v7
            case -9223372036854775808:
                p0 = p0
            case -292:
                for fi8 := 0; fi8 < 18; fi8++ {
                    var v9 uint64 = 0
                    _ = fi8
                    _ = v9
                }
            default:
                var v10 float32 = 1.5
                _ = v10
            }
            var v11 int32 = 2147483647
            _ = v11
        }
        p0 -= (p0 | p0)
        _ = fi5
        _ = v6
    }
    var v12 int64 = 1
    _ = p0
    _ = v12
    _ = v4
    return false
}
func genFn1() string {
    var v13 int8 = 111
    var v14 uint32 = 0
    v15 := []int{}
    _ = v13
    _ = v14
    _ = v15
    return "a"
}

func main() {
    var v1 int = 127
    var v2 uint64 = 42
    var v3 uint = 65535
    var v16 uint16 = 32767
    v17 := []int{2147483647}
    if len(v17) > 0 {
        v17[0] = v1
    }
    if false {
        if false {
            var v18 uint32 = 695
            v19 := GenS1{F0: (v3 >> ((v3) & 63))}
            _ = v18
            _ = v19
        } else {
            v20 := GenS2{F0: 42}
            if ((!true) || false) {
                v17 = append(v17, (v1 * 1))
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                switch "\n\t" {
                case "hello":
                    v1 = 4294967295
                case "\n\t":
                    v1 = (v1 >> ((v3) & 63))
                }
            }
            _ = v20
        }
        v3 = (v3 % (((v3 & (v3 / (((^(v3))) | 1)))) | 1))
    }
    var v21 string = "\n\t"
    for fi22 := 0; fi22 < 3; fi22++ {
        if (!true) {
            for fi23 := 0; fi23 < 10; fi23++ {
                v16 += v16
                _ = fi23
            }
            for fi24 := 0; fi24 < 6; fi24++ {
                v1 -= -888
                v25 := GenS3{F0: (v3 - 42)}
                _ = fi24
                _ = v25
            }
            if (!(0.0 == (1.5 / 3.14159))) {
                var v26 uint8 = uint8(v3)
                v27 := GenS2{F0: uint8(v16)}
                fmt.Println(v2, v1)
                _ = v26
                _ = v27
            }
        } else {
            switch (-(fi22)) {
            case 9223372036854775807:
                var v28 int64 = 42
                _ = v28
            }
            v1 += fi22
            fmt.Println(v16, v2, v2)
        }
        _ = len(v17)
        _ = fi22
    }
    v29 := GenS3{F0: v3}
    var v30 int = 42
    v3 += 809
    var v31 uint8 = 1
    v21 = (v21 + "\n\t")
    v31 = (v31 * v31)
    v21 = "a"
    if (42 <= 128) {
        var v32 bool = (!(float32(float64(1.5)) <= (1.5 - ((3.14159 + (1e10 * 1.5)) + (float32(3.14159) * 3.14159)))))
        v33 := GenS4{F0: (!(v3 < (v3 >> (((v3 + uint(v31))) & 63)))), F1: (v1 - 32768), F2: (v21 + v21)}
        _ = v32
        _ = v33
    }
    if true {
        var v34 uint64 = (v2 ^ v2)
        var v35 uint32 = 616
        _ = v34
        _ = v35
    }
    if len(v17) > 0 {
        v17[0] = -2147483648
    }
    v36 := GenS3{F0: 42}
    if true {
        var v37 uint32 = 42
        v3 = v3
        v37 -= (v37 ^ v37)
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
}
