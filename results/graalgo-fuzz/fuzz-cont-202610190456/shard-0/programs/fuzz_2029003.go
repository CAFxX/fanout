package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint64 }

func main() {
    var v1 uint = 65535
    var v2 int = -9223372036854775808
    var v3 uint = (v1 + (v1 ^ v1))
    var v4 string = ("" + "")
    var v5 uint16 = 223
    var v6 string = (("foo bar" + "a") + "")
    var v7 uint32 = 0
    v8 := GenS1{F0: -1, F1: (v4 + v6), F2: 42}
    var v9 int32 = -1
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS2{F0: (v5 << (((^(v1))) & 63)), F1: 56, F2: 0}
        _ = fi10
        _ = v11
    }
    v6 = v4
    if false {
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := []int{v2}
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 float32 = 1.5
                fmt.Println(v7)
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                var v17 bool = ((255 != (v3 / ((v3) | 1))) && (0.1 < (0.0 * 3.14159)))
                v6 = (v4 + v6)
                _ = fi16
                _ = v17
            }
            _ = fi12
            _ = v13
        }
        v18 := GenS3{F0: 42}
        for fi19 := 0; fi19 < 5; fi19++ {
            switch fi19 {
            case -659:
                v9 -= -1
            case -9223372036854775808:
                v20 := []int{32767, v2, fi19}
                _ = v20
            }
            v21 := []int{183, (fi19 * fi19), fi19, 1}
            for fi22 := 0; fi22 < 4; fi22++ {
                fmt.Println(fi22)
                _ = fi22
            }
            _ = fi19
            _ = v21
        }
        _ = v18
    } else {
        v23 := GenS4{F0: (v2 % (((v2 ^ v2)) | 1)), F1: -1}
        for fi24 := 0; fi24 < 18; fi24++ {
            for fi25 := 0; fi25 < 7; fi25++ {
                fmt.Println(v4, v2, v1)
                var v26 uint64 = 1
                _ = fi25
                _ = v26
            }
            var v27 bool = (false && false)
            _ = fi24
            _ = v27
        }
        _ = v23
    }
    var v28 uint32 = uint32(v3)
    v29 := GenS1{F0: 0, F1: "a", F2: -80}
    v30 := []int{int(v5), v2, 4294967295}
    v31 := GenS5{F0: (v5 * (v5 >> (((v1 + (v1 + (v3 ^ 332)))) & 63))), F1: 103, F2: 654}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
