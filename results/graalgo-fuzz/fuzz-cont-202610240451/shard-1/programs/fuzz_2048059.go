package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8; F1 int; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 int64; F2 bool }
type GenS5 struct { F0 int32; F1 uint64; F2 int8 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 uint32; F1 int8; F2 string }
type GenS8 struct { F0 bool }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]uint32)
    var v3 uint16 = 741
    if (2 < 4294967295) {
        v1 += v1
    }
    fmt.Println(v3, v1)
    for fi4 := 0; fi4 < 12; fi4++ {
        if ((v1 | (v1 - (-(v1)))) <= 1) {
            var v5 uint = 896
            var v6 string = ""
            v7 := GenS1{F0: 1, F1: (v5 - v5), F2: v3}
            _ = v5
            _ = v6
            _ = v7
        }
        fmt.Println(fi4, fi4)
        _ = fi4
    }
    v8 := GenS2{F0: (v1 / ((v1) | 1)), F1: 0, F2: 766}
    fmt.Println(v1)
    switch 32768 {
    case 491:
        v9 := []int{4294967295, -2147483648, 42}
        _ = v9
    default:
        fmt.Println(v3, v1)
    }
    var v10 uint32 = 2147483648
    switch 467 {
    case 9223372036854775807:
        delete(v2, "")
    default:
        v11 := GenS1{F0: 433, F1: 224, F2: v3}
        _ = v11
    }
    v12 := GenS2{F0: -128, F1: -2147483648, F2: 666}
    delete(v2, "")
    fmt.Println(v1)
    v13 := GenS1{F0: 4294967295, F1: 2, F2: 32767}
    if ((!true) || (((false || (!true)) && (v3 > v3)) && true)) {
        var v14 int8 = (v1 - (v1 * (v1 | v1)))
        v15 := GenS3{F0: v1}
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS4{F0: "hello", F1: 640, F2: true}
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    }
    var v18 float32 = 1.5
    switch "" {
    case "\n\t":
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(v1)
            switch "" {
            case "hello":
                if true {
                    v20 := GenS3{F0: (v1 % ((100) | 1))}
                    var v21 uint64 = 480
                    switch (("foo bar" + "") + ("\n\t" + "foo bar")) {
                    case "a":
                        v22 := []int{int(v3), (fi19 % ((-439) | 1)), -9223372036854775808}
                        _ = v22
                    default:
                        v23 := GenS5{F0: 365, F1: (v21 ^ (v21 << ((42) & 63))), F2: -128}
                        _ = v23
                    }
                    _ = v20
                    _ = v21
                }
            default:
                var v24 int16 = -537
                _ = v24
            }
            _ = fi19
        }
    case "hello":
        for fi25 := 0; fi25 < 14; fi25++ {
            switch "" {
            case "a":
                for fi26 := 0; fi26 < 3; fi26++ {
                    switch fi25 {
                    case 42:
                        v27 := v2["hello"]
                        _ = v27
                    case 0:
                        v1 = 123
                    }
                    v3 -= 32767
                    _ = fi26
                }
            case "":
                fmt.Println(fi25)
            case "foo bar":
                switch fi25 {
                case 1:
                    v28 := GenS6{F0: v1, F1: (v1 << ((599) & 63))}
                    _ = v28
                case 42:
                    v29 := []int{fi25}
                    _ = v29
                case 128:
                    var v30 float32 = (1e10 * 3.14159)
                    _ = v30
                default:
                    if ((v10 % (((v10 | 58)) | 1)) == v10) {
                        var v31 int8 = v1
                        v2["hello"] = 2
                        if false {
                            v32 := GenS7{F0: 0, F1: (-(v1)), F2: "hello"}
                            _ = v32
                        }
                        _ = v31
                    }
                }
            }
            v33 := GenS5{F0: 112, F1: 1, F2: 42}
            _ = fi25
            _ = v33
        }
    case "foo bar":
        switch "foo bar" {
        case "hello":
            var v34 int16 = -1
            _ = v34
        case "\n\t":
            if (-1 != -9223372036854775808) {
                v35 := GenS8{F0: false}
                _ = v35
            }
        }
    default:
        for fi36 := 0; fi36 < 12; fi36++ {
            for fi37 := 0; fi37 < 14; fi37++ {
                var v38 int8 = (v1 % ((42) | 1))
                var v39 float64 = 1.5
                _ = fi37
                _ = v38
                _ = v39
            }
            var v40 int32 = 1
            _ = fi36
            _ = v40
        }
    }
    var v41 int32 = -808
    var v42 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v8
}
