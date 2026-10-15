package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int8; F1 uint8; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 string; F2 int }
func genFn0(p0 uint32, p1 int16) uint8 {
    for fi6 := 0; fi6 < 12; fi6++ {
        p0 += p0
        switch "a" {
        case "a":
            switch fi6 {
            case 32768:
                v7 := GenS1{F0: 1}
                _ = v7
            case 127:
                p1 = 42
            case 32767:
                p0 += (p0 - 2147483648)
            }
        default:
            v8 := GenS1{F0: 65535}
            _ = v8
        }
        _ = fi6
    }
    var v9 int64 = -405
    _ = p0
    _ = p1
    _ = v9
    return 172
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 uint64 = 4294967295
    v4 := []string{"", ("a" + ("" + ("foo bar" + "hello")))}
    v5 := []int{}
    var v10 bool = (386 != 32767)
    if (false || true) {
        for fi11 := 0; fi11 < 11; fi11++ {
            v10 = ((-2.25 > float32(3.14159)) || v10)
            _ = fi11
        }
        if (0 == 178) {
            v12 := GenS1{F0: 802}
            v13 := GenS2{F0: 640, F1: (false || (!true))}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(fi14, v3)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi14
            }
        }
        if (((v3 ^ (v3 - 1)) < v3) && ((2 <= 32767) && (genFn0(1, 32767) > 2))) {
            v15 := GenS3{F0: -1, F1: 180, F2: -353}
            _ = v15
        } else {
            v5 = append(v5, 0)
            _ = len(v1)
            _ = len(v4)
        }
    } else {
        var v16 float64 = 1.5
        _ = v16
    }
    if v10 {
        var v17 bool = ("hello" < "")
        _ = v17
    } else {
        for fi18 := 0; fi18 < 14; fi18++ {
            v2[fi18] = (("" + ("a" + "foo bar")) + "a")
            if false {
                fmt.Println(fi18, v3, v3)
                var v19 uint8 = genFn0(1, -316)
                v10 = v10
                _ = v19
            } else {
                var v20 int16 = 824
                _ = v20
            }
            _ = fi18
        }
        v10 = false
    }
    v10 = (v10 && v10)
    v21 := GenS3{F0: 1, F1: 0, F2: 895}
    var v22 float32 = 1.5
    v23 := GenS4{F0: 128}
    var v24 int64 = 42
    v25 := GenS5{F0: -815, F1: "a", F2: 32767}
    for fi26 := 0; fi26 < 7; fi26++ {
        fmt.Println(v3, v24)
        var v27 int = 127
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
