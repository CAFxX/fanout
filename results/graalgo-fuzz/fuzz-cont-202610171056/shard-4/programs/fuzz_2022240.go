package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 uint64 }
func genFn0() uint32 {
    v4 := []int{663, 4294967295, 32767, -892}
    var v5 bool = (1 != 211)
    v6 := GenS1{F0: 1}
    fmt.Println(v5)
    _ = v4
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    var v1 int16 = 577
    v2 := make(map[int]int)
    var v3 int8 = 0
    var v7 float64 = -2.25
    v8 := v2[2147483647]
    if false {
        _ = len(v2)
    }
    switch (v8 >> ((18446744073709551615) & 63)) {
    case -2147483648:
        v9 := GenS1{F0: 4294967295}
        _ = v9
    case 1:
        v10 := GenS1{F0: 446}
        _ = v10
    }
    switch 32768 {
    case 4294967295:
        v11 := v2[4294967295]
        _ = v11
    case 237:
        if false {
            v1 -= v1
        }
    case -58:
        for fi12 := 0; fi12 < 1; fi12++ {
            v8 -= v8
            if (489 == 255) {
                fmt.Println(v8)
                fmt.Println(v8, v7)
                for fi13 := 0; fi13 < 14; fi13++ {
                    switch "" {
                    case "a":
                        var v14 uint8 = 129
                        _ = v14
                    case "hello":
                        switch fi13 {
                        case 32767:
                            fmt.Println(fi12)
                        case -976:
                            v15 := GenS2{F0: 1, F1: "\n\t"}
                            _ = v15
                        }
                    case "\n\t":
                        if true {
                            fmt.Println(fi13, v3, v3)
                        }
                    }
                    if (!true) {
                        switch v8 {
                        case 2147483647:
                            var v16 uint16 = 255
                            _ = v16
                        case 4294967295:
                            v7 -= v7
                            var v17 uint64 = 0
                            var v18 bool = ("hello" >= "hello")
                            _ = v17
                            _ = v18
                        }
                        switch v8 {
                        case 856:
                            v19 := GenS2{F0: 65535, F1: ("a" + ((("foo bar" + "\n\t") + "") + "hello"))}
                            _ = v19
                        }
                        v20 := GenS3{F0: true, F1: 461}
                        _ = v20
                    }
                    _ = fi13
                }
            } else {
                switch "\n\t" {
                case "foo bar":
                    var v21 uint64 = 0
                    _ = v21
                case "":
                    for fi22 := 0; fi22 < 2; fi22++ {
                        fmt.Println(v3, v8, v8)
                        if true {
                            var v23 int32 = 0
                            _ = v23
                        } else {
                            var v24 uint16 = 2
                            v25 := v2[fi12]
                            v8 += (fi12 & 2147483647)
                            _ = v24
                            _ = v25
                        }
                        v3 = v3
                        fmt.Println(fi22, fi22, v7)
                        switch v8 {
                        case -9223372036854775808:
                            v26 := GenS1{F0: 2147483648}
                            _ = v26
                        case 0:
                            var v27 int8 = v3
                            _ = v27
                        }
                        v28 := []int{(v8 - (fi22 ^ (v8 << ((65535) & 63))))}
                        _ = fi22
                        _ = v28
                    }
                case "hello":
                    v29 := GenS2{F0: 2, F1: ("a" + "\n\t")}
                    _ = v29
                }
                if false {
                    switch (v8 ^ 9223372036854775807) {
                    case 770:
                        for fi30 := 0; fi30 < 6; fi30++ {
                            fmt.Println(v8, v7, fi30)
                            v2[(-(v8))] = fi12
                            _ = fi30
                        }
                    case 9223372036854775807:
                        v3 += v3
                    case 1:
                        if false {
                            v31 := v2[v8]
                            delete(v2, -1)
                            v32 := []int{fi12}
                            _ = v31
                            _ = v32
                        }
                    default:
                        var v33 uint16 = 941
                        _ = v33
                    }
                    switch fi12 {
                    case 32767:
                        v34 := GenS4{F0: int(v3)}
                        _ = v34
                    case 32768:
                        v8 += (fi12 * 32768)
                    case 9223372036854775807:
                        if (467 == -953) {
                            var v35 int32 = -2147483648
                            v36 := v2[(fi12 >> ((761) & 63))]
                            _ = v35
                            _ = v36
                        } else {
                            v1 = v1
                        }
                    default:
                        var v37 string = "a"
                        _ = v37
                    }
                }
                if false {
                    v38 := GenS5{F0: 42, F1: 18446744073709551615}
                    for fi39 := 0; fi39 < 17; fi39++ {
                        if false {
                            var v40 int = fi12
                            var v41 uint8 = 0
                            _ = v40
                            _ = v41
                        }
                        v42 := v2[172]
                        _ = fi39
                        _ = v42
                    }
                    _ = v38
                } else {
                    var v43 int = (v8 << ((1) & 63))
                    var v44 float32 = (1e10 + ((1e10 - (0.1 * ((float32((v7 + 1.5)) / float32(float64((1e10 - 1.5)))) * (3.14159 + ((1.5 - 1.5) / -2.25))))) / (1.5 / 3.14159)))
                    if ((v44 / -2.25) > 1e10) {
                        if true {
                            v7 = (0.1 / v7)
                        } else {
                            v1 = 1000
                            v7 = (v7 + (v7 * float64(v44)))
                        }
                    }
                    _ = v43
                    _ = v44
                }
            }
            _ = fi12
        }
    }
    v45 := GenS1{F0: 4294967295}
    for fi46 := 0; fi46 < 11; fi46++ {
        if (!false) {
            var v47 uint64 = 1
            v8 = (fi46 >> ((29) & 63))
            v7 = v7
            _ = v47
        }
        v48 := []int{v8, fi46, v8, v8}
        _ = fi46
        _ = v48
    }
    fmt.Println(v3, v7, v8)
    var v49 uint = 18446744073709551615
    fmt.Println(v3)
    v7 += float64(-2.25)
    if ((v8 != 32767) || true) {
        v50 := GenS5{F0: 2, F1: 0}
        var v51 float64 = v7
        _ = v50
        _ = v51
    }
    var v52 uint = (v49 + (v49 + 65535))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v52:", v52)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v49
    _ = v52
    _ = v7
    _ = v8
}
