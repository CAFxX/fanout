package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 bool; F1 int16 }
func genFn0() int32 {
    switch -895 {
    case 1:
        switch ("hello" + "hello") {
        case "hello":
            var v6 float32 = (3.14159 + 1e10)
            _ = v6
        case "a":
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 uint16 = 65535
                _ = fi7
                _ = v8
            }
        case "":
            var v9 int64 = 741
            _ = v9
        }
    case 42:
        switch -699 {
        case 128:
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v10 uint16 = 2
                _ = v10
            case 4294967295:
                var v11 uint64 = 2
                _ = v11
            default:
                for fi12 := 0; fi12 < 10; fi12++ {
                    if ((1e10 / (float32((1.5 - ((0.1 - 3.14159) / 1.5))) / 1.0)) < 1.5) {
                        var v13 int = (fi12 - 127)
                        _ = v13
                    }
                    v14 := []int{fi12, (^(fi12))}
                    v15 := GenS1{F0: -32768}
                    _ = fi12
                    _ = v14
                    _ = v15
                }
            }
        default:
            var v16 uint16 = 2
            _ = v16
        }
    }
    var v17 float32 = 0.1
    if (-1 != 32767) {
        v17 -= v17
    } else {
        switch -623 {
        case 922:
            for fi18 := 0; fi18 < 1; fi18++ {
                if (!true) {
                    v17 -= float32(float64(v17))
                    var v19 float32 = float32(1.5)
                    _ = v19
                }
                v20 := GenS2{F0: 505, F1: 121, F2: 4294967295}
                var v21 float64 = (0.1 / -2.25)
                _ = fi18
                _ = v20
                _ = v21
            }
        case -9223372036854775808:
            fmt.Println(v17)
        case 0:
            for fi22 := 0; fi22 < 8; fi22++ {
                for fi23 := 0; fi23 < 18; fi23++ {
                    v17 -= float32((1.5 / (float64(0.1) / 3.14159)))
                    _ = fi23
                }
                _ = fi22
            }
        }
    }
    if (2147483648 < 2) {
        fmt.Println(v17)
        v17 = v17
        for fi24 := 0; fi24 < 18; fi24++ {
            fmt.Println(fi24, v17)
            _ = fi24
        }
    }
    var v25 int8 = 42
    _ = v17
    _ = v25
    return -577
}

func main() {
    var v1 string = "hello"
    var v2 int = 32767
    var v3 int64 = 259
    v4 := []uint32{337, 4294967295, 0, 385}
    var v5 float32 = float32(1e10)
    for fi26 := 0; fi26 < 15; fi26++ {
        v5 = (v5 * (v5 + v5))
        v27 := GenS2{F0: 255, F1: 1, F2: 2147483647}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        var v29 int = 42
        if false {
            if (true && ((v3 / ((-807) | 1)) < v3)) {
                var v30 int = (fi28 >> ((255) & 63))
                _ = v30
            }
            if len(v4) > 0 {
                v4[0] = 1
            }
            var v31 uint64 = 1
            _ = v31
        } else {
            var v32 float64 = 3.14159
            fmt.Println(v1, v2)
            v29 += (v2 & v29)
            _ = v32
        }
        v29 = v29
        _ = fi28
        _ = v29
    }
    v1 = "foo bar"
    v33 := GenS1{F0: -963}
    for fi34 := 0; fi34 < 8; fi34++ {
        fmt.Println(v5, v2, v1)
        var v35 int32 = 60
        _ = fi34
        _ = v35
    }
    var v36 float64 = 0.1
    if true {
        _ = len(v4)
    }
    var v37 bool = true
    var v38 uint32 = 382
    v5 += (v5 / -2.25)
    var v39 int8 = 1
    if (3.14159 == v36) {
        v40 := GenS2{F0: 2, F1: (v38 & v38), F2: v38}
        v41 := GenS2{F0: 1, F1: (v38 * v38), F2: 42}
        if len(v4) > 0 {
            v4[0] = v38
        }
        _ = v40
        _ = v41
    } else {
        if (false || v37) {
            switch (-(v2)) {
            case -9223372036854775808:
                var v42 uint16 = 442
                _ = v42
            case -689:
                switch "" {
                case "":
                    v43 := GenS1{F0: 32767}
                    _ = v43
                case "foo bar":
                    v44 := GenS1{F0: 565}
                    _ = v44
                case "\n\t":
                    var v45 bool = (4294967295 == uint(v38))
                    _ = v45
                }
            default:
                for fi46 := 0; fi46 < 8; fi46++ {
                    for fi47 := 0; fi47 < 10; fi47++ {
                        v2 -= fi47
                        _ = fi47
                    }
                    v4 = append(v4, (v38 << ((2) & 63)))
                    _ = fi46
                }
            }
            v5 -= v5
            v48 := GenS3{F0: true, F1: 1}
            _ = v48
        } else {
            var v49 int = (v2 + v2)
            for fi50 := 0; fi50 < 13; fi50++ {
                var v51 int32 = 42
                _ = fi50
                _ = v51
            }
            _ = v49
        }
        fmt.Println(v39)
    }
    v4 = append(v4, v38)
    v5 = float32(((v36 / (0.1 - float64(v5))) * 1.5))
    _ = len(v4)
    v38 += (^(v38))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
