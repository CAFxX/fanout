package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int8; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 string; F2 int8 }
type GenS6 struct { F0 bool; F1 int8; F2 int16 }
func genFn0(p0 uint64) uint8 {
    switch 32767 {
    case 32768:
        var v7 int8 = 0
        _ = v7
    case 9223372036854775807:
        var v8 uint64 = 42
        _ = v8
    case -770:
        var v9 string = ((("a" + "foo bar") + "\n\t") + (("a" + "\n\t") + (("a" + "\n\t") + "a")))
        _ = v9
    default:
        v10 := []int{1, 461, 4294967295, 1}
        _ = v10
    }
    p0 = 997
    v11 := []int{42, -9223372036854775808, 32767, 32767}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v11 = append(v11, 116)
    _ = p0
    _ = v11
    return 1
}

func main() {
    var v1 int8 = 1
    var v2 float32 = (float32(((3.14159 * 0.1) * 1.5)) - 0.0)
    var v3 uint64 = 42
    v4 := make(map[int]bool)
    var v5 bool = (0.0 > -2.25)
    var v6 float32 = 0.0
    for fi12 := 0; fi12 < 2; fi12++ {
        switch 0 {
        case 1:
            var v13 int64 = 474
            _ = v13
        }
        fmt.Println(fi12, fi12)
        v14 := v4[4294967295]
        _ = fi12
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := []int{int(v3), fi15, (-(fi15))}
        _ = fi15
        _ = v16
    }
    if true {
        v17 := GenS1{F0: 1}
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        if false {
            fmt.Println(fi18, v6)
            var v19 uint64 = (v3 << ((1) & 63))
            v6 = v2
            _ = v19
        } else {
            var v20 float32 = 0.0
            _ = v20
        }
        _ = fi18
    }
    if (298 != 2147483647) {
        fmt.Println(v1, v3, v1)
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS2{F0: 42}
            if v5 {
                if ((fi21 - fi21) >= (fi21 + -421)) {
                    v6 = -2.25
                    var v23 int = -843
                    _ = v23
                } else {
                    if v5 {
                        fmt.Println(fi21, v3)
                        v24 := GenS1{F0: -1}
                        _ = v24
                    }
                }
                v25 := GenS3{F0: 0, F1: (v1 % (((v1 + v1)) | 1)), F2: v5}
                v1 = v1
                _ = v25
            } else {
                fmt.Println(v2, v5, v5)
            }
            _ = fi21
            _ = v22
        }
    }
    var v26 int8 = v1
    if (711 >= 1) {
        var v27 int32 = -940
        if true {
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS1{F0: (v27 - v27)}
                var v30 int16 = int16(v1)
                v31 := []int{fi28, (fi28 - fi28), fi28}
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
        }
        v32 := GenS4{F0: "hello"}
        _ = v27
        _ = v32
    }
    v33 := GenS5{F0: v3, F1: "foo bar", F2: 35}
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 int8 = 100
        _ = fi34
        _ = v35
    }
    v36 := v4[9223372036854775807]
    switch "a" {
    case "hello":
        if (v3 == v3) {
            var v37 int8 = (v26 % ((1) | 1))
            var v38 bool = ("a" <= "foo bar")
            _ = v37
            _ = v38
        } else {
            fmt.Println(v1)
        }
    case "":
        var v39 float32 = v2
        _ = v39
    default:
        var v40 float64 = (0.1 + (0.1 - (0.0 * -2.25)))
        _ = v40
    }
    _ = len(v4)
    for fi41 := 0; fi41 < 2; fi41++ {
        switch "hello" {
        case "":
            if v5 {
                fmt.Println(v3)
            }
        case "a":
            if (1 <= 479) {
                fmt.Println(v3)
            } else {
                var v42 int32 = -2147483648
                for fi43 := 0; fi43 < 1; fi43++ {
                    v2 -= 1.5
                    if v5 {
                        var v44 bool = ((v36 || v36) && false)
                        fmt.Println(v6, fi43, v2)
                        v42 -= v42
                        var v45 float64 = float64(-2.25)
                        _ = v44
                        _ = v45
                    } else {
                        var v46 uint16 = 65535
                        v26 += (v1 - 42)
                        v26 = 25
                        _ = v46
                    }
                    v47 := []int{(fi41 + 42), fi43, (fi43 << ((54) & 63)), fi43}
                    _ = fi43
                    _ = v47
                }
                v26 -= int8(v3)
                _ = v42
            }
        case "\n\t":
            _ = len(v4)
        default:
            fmt.Println(v2)
        }
        fmt.Println(v3, v36, v36)
        _ = fi41
    }
    v48 := GenS6{F0: v36, F1: 15, F2: 1000}
    var v49 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
