package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 uint }
func genFn0(p0 int64) bool {
    var v5 int16 = 1
    v5 = -32768
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(v5, v5)
            v8 := []int{9223372036854775807, (fi7 | fi6), fi7, (-(fi6))}
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: (v5 & v5), F1: (v5 | v5)}
        for fi10 := 0; fi10 < 3; fi10++ {
            switch (-(fi6)) {
            case 0:
                var v11 float64 = 0.1
                var v12 int = fi10
                if false {
                    switch (fi6 & fi10) {
                    case -760:
                        fmt.Println(v11, p0, fi10)
                    case 0:
                        var v13 int8 = 100
                        _ = v13
                    case -2147483648:
                        var v14 int = fi10
                        _ = v14
                    default:
                        v12 += (^(fi6))
                    }
                    v11 += -2.25
                    v15 := []int{}
                    _ = v15
                } else {
                    switch "foo bar" {
                    case "":
                        v16 := []int{(fi10 >> ((65535) & 63))}
                        _ = v16
                    }
                    switch (fi6 / (((^(fi10))) | 1)) {
                    case -146:
                        v17 := GenS2{F0: p0, F1: 0, F2: 50}
                        _ = v17
                    }
                }
                var v18 int8 = 0
                _ = v11
                _ = v12
                _ = v18
            }
            if false {
                var v19 string = "hello"
                var v20 int = (fi6 ^ -2147483648)
                _ = v19
                _ = v20
            }
            v21 := GenS2{F0: (p0 + 0), F1: 894, F2: 255}
            _ = fi10
            _ = v21
        }
        _ = fi6
        _ = v9
    }
    var v22 float64 = 1.5
    if true {
        if false {
            if true {
                var v23 string = "\n\t"
                _ = v23
            }
        } else {
            var v24 float32 = float32((3.14159 / v22))
            v25 := GenS2{F0: (^(p0)), F1: 65535, F2: 255}
            if (-1 > -374) {
                fmt.Println(v22, v22)
                v26 := GenS3{F0: 4294967295, F1: false, F2: 65535}
                _ = v26
            }
            _ = v24
            _ = v25
        }
        fmt.Println(p0, p0, v22)
    } else {
        p0 -= -1
    }
    _ = p0
    _ = v22
    _ = v5
    return (-9223372036854775808 >= (p0 ^ (p0 % ((-84) | 1))))
}

func main() {
    var v1 float32 = ((1e10 + 3.14159) - 0.1)
    var v2 uint32 = 774
    var v3 uint32 = (v2 / ((65535) | 1))
    var v4 int = 0
    var v27 int = 128
    var v28 uint = 255
    switch ("foo bar" + "a") {
    case "foo bar":
        v27 = v4
    default:
        var v29 float64 = 1e10
        _ = v29
    }
    var v30 float32 = v1
    v4 += (v27 << (((v28 % ((uint(v2)) | 1))) & 63))
    for fi31 := 0; fi31 < 6; fi31++ {
        for fi32 := 0; fi32 < 5; fi32++ {
            v33 := GenS3{F0: 9223372036854775807, F1: false, F2: 255}
            v4 = (fi32 - (fi32 * -281))
            v34 := []int{}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    if false {
        v27 = v27
    } else {
        switch ("\n\t" + "a") {
        case "\n\t":
            switch "a" {
            case "":
                var v35 int16 = 42
                _ = v35
            case "hello":
                v4 = v27
            case "foo bar":
                var v36 uint32 = 13
                _ = v36
            }
        case "hello":
            var v37 uint64 = 4294967295
            _ = v37
        }
        var v38 float64 = 0.1
        fmt.Println(v27, v38)
        _ = v38
    }
    v2 -= uint32(v28)
    v30 += float32((-2.25 + 1e10))
    var v39 int32 = -657
    var v40 string = ""
    switch v4 {
    case -561:
        v41 := []int{86}
        _ = v41
    case 2147483647:
        var v42 int = v27
        _ = v42
    }
    switch (v27 >> (((v28 - 0)) & 63)) {
    case 262:
        v43 := []int{32767}
        _ = v43
    case -649:
        for fi44 := 0; fi44 < 12; fi44++ {
            v1 += 1.5
            _ = fi44
        }
    default:
        var v45 int64 = 0
        _ = v45
    }
    v40 = "a"
    var v46 int = (v4 * v27)
    v46 = (v46 << ((255) & 63))
    var v47 uint32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v46
    _ = v47
}
