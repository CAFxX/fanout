package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint32; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int16 = 42
    var v2 uint32 = 2
    var v3 string = ("" + "")
    v4 := make(map[string]uint32)
    v5 := GenS1{F0: 0}
    if true {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 int = (fi6 % (((fi6 | -1)) | 1))
            _ = fi6
            _ = v7
        }
        v2 -= v2
        v8 := []int{32767, 4294967295}
        _ = v8
    } else {
        if (true && false) {
            fmt.Println(v2)
            var v9 uint8 = 0
            _ = v9
        } else {
            var v10 bool = (!(531 != 65535))
            _ = len(v4)
            _ = v10
        }
        if false {
            var v11 int16 = -83
            v3 = v3
            var v12 int8 = -54
            _ = v11
            _ = v12
        } else {
            var v13 bool = false
            v2 = v2
            v14 := GenS2{F0: -2147483648}
            _ = v13
            _ = v14
        }
    }
    v1 -= 1000
    v15 := []int{1}
    if (v1 <= (v1 & (v1 % (((v1 << ((4294967295) & 63))) | 1)))) {
        v1 = 461
        var v16 int32 = 0
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint8 = 128
            var v19 uint64 = 468
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    }
    var v20 string = v3
    var v21 uint = 18446744073709551615
    fmt.Println(v1, v1)
    v20 = v20
    fmt.Println(v20)
    var v22 uint64 = 42
    if true {
        if true {
            v21 -= v21
        } else {
            for fi23 := 0; fi23 < 12; fi23++ {
                v24 := GenS1{F0: v21}
                fmt.Println(v21, v20)
                _ = fi23
                _ = v24
            }
        }
        v25 := GenS2{F0: -2147483648}
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS3{F0: (^(v21)), F1: 358, F2: fi26}
            switch fi26 {
            case -9223372036854775808:
                fmt.Println(v3)
            case 1:
                v21 = 4294967295
            case 249:
                for fi28 := 0; fi28 < 13; fi28++ {
                    if (-976 < -9223372036854775808) {
                        v20 = (v20 + "")
                        var v29 uint16 = 42
                        _ = v29
                    }
                    _ = fi28
                }
            }
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    var v30 int = -853
    var v31 int64 = 42
    if (v1 == v1) {
        switch (v30 << (((v21 << ((v21) & 63))) & 63)) {
        case 9223372036854775807:
            var v32 bool = true
            _ = v32
        case -428:
            v33 := GenS4{F0: 460}
            _ = v33
        }
    }
    if ((v22 / ((2) | 1)) != (^(v22))) {
        v20 = (v20 + "\n\t")
        v34 := GenS4{F0: (v1 + -334)}
        v20 = v20
        _ = v34
    }
    if false {
        for fi35 := 0; fi35 < 2; fi35++ {
            if false {
                var v36 int8 = -128
                _ = v36
            } else {
                v31 -= 9223372036854775807
            }
            _ = fi35
        }
        v37 := GenS5{F0: v1}
        _ = v37
    }
    var v38 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
    _ = v5
}
