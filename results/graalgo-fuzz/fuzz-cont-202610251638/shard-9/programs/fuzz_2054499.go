package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint8 }
func genFn0(p0 int8) uint16 {
    fmt.Println(p0)
    var v7 uint32 = 65535
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{2147483647, (fi9 & -551)}
            var v11 uint8 = 2
            p0 += (^(p0))
            _ = fi9
            _ = v10
            _ = v11
        }
        switch fi8 {
        case -2147483648:
            v12 := GenS1{F0: fi8, F1: 134}
            _ = v12
        default:
            var v13 uint8 = 0
            _ = v13
        }
        var v14 int = 846
        _ = fi8
        _ = v14
    }
    _ = p0
    _ = v7
    return 477
}
func genFn1() int16 {
    var v15 string = "a"
    if false {
        v16 := []int{128}
        _ = v16
    }
    _ = v15
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 209
    var v3 int16 = 42
    v4 := make(map[uint32]bool)
    var v5 float64 = 0.1
    var v6 float64 = v5
    v17 := GenS1{F0: 9223372036854775807, F1: 4294967295}
    fmt.Println(v2, v6, v6)
    switch "a" {
    case "hello":
        v18 := GenS2{F0: 821, F1: 4294967295, F2: "hello"}
        _ = v18
    case "":
        for fi19 := 0; fi19 < 6; fi19++ {
            v5 += float64(3.14159)
            _ = fi19
        }
    case "\n\t":
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 uint = 255
            var v22 uint = (v21 >> ((255) & 63))
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    if (!true) {
        v23 := GenS2{F0: 380, F1: 2147483647, F2: ("hello" + "")}
        if true {
            switch "" {
            case "a":
                var v24 bool = false
                _ = v24
            case "hello":
                for fi25 := 0; fi25 < 4; fi25++ {
                    v6 = ((-2.25 + 1e10) * float64(float32(v6)))
                    _ = fi25
                }
            }
        } else {
            var v26 string = "hello"
            var v27 float64 = 0.0
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 16; fi28++ {
            fmt.Println(fi28, v2, v5)
            v5 = 0.0
            _ = fi28
        }
        _ = v23
    }
    var v29 int64 = 9223372036854775807
    if false {
        v30 := []int{}
        fmt.Println(v29, v6, v6)
        var v31 float32 = 1.5
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 float32 = 3.14159
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        v35 := GenS1{F0: fi34, F1: 42}
        _ = fi34
        _ = v35
    }
    if true {
        fmt.Println(v6)
    }
    fmt.Println(v29, v6)
    var v36 uint8 = v2
    var v37 uint64 = 2
    var v38 uint64 = 4294967295
    for fi39 := 0; fi39 < 9; fi39++ {
        switch -114 {
        case -536:
            v40 := []int{127}
            _ = v40
        case 347:
            if (false && true) {
                v41 := []int{}
                if true {
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                } else {
                    v4[684] = true
                    var v42 uint = 474
                    _ = v42
                }
                _ = v41
            } else {
                v43 := GenS3{F0: (^(v37)), F1: 146, F2: (v2 | v2)}
                _ = v43
            }
        default:
            v3 = v3
        }
        var v44 uint = 2
        var v45 int = int(v2)
        _ = fi39
        _ = v44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
