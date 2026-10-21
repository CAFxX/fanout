package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int; F1 int32 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]uint32)
    var v3 string = "a"
    var v4 float64 = 1e10
    var v5 int8 = -22
    if (!(!(!(!(!true))))) {
        var v6 uint32 = 1
        if (-895 == 2147483647) {
            v6 += v6
        }
        _ = v6
    }
    v7 := v2[412]
    v8 := GenS1{F0: int32(v5), F1: ((!((2 < 2) && (255 > 1))) && (2147483647 < v7))}
    fmt.Println(v1, v5)
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            v3 = "a"
            _ = fi10
        }
        _ = fi9
    }
    v2[(^(v7))] = v7
    switch ("hello" + v3) {
    case "a":
        v11 := []int{2147483647, 32767, 384, -90}
        _ = v11
    case "\n\t":
        var v12 uint8 = 127
        _ = v12
    case "foo bar":
        if false {
            v1 += v1
        }
    default:
        var v13 int8 = v5
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int = (fi14 % (((fi14 % ((32768) | 1))) | 1))
        _ = fi14
        _ = v15
    }
    var v16 float64 = (v4 * 0.1)
    if false {
        for fi17 := 0; fi17 < 19; fi17++ {
            switch "a" {
            case "hello":
                v16 += ((v16 / float64(3.14159)) - (v16 / ((v4 * (3.14159 - v16)) - float64((1.5 / ((0.0 - float32(v4)) * 3.14159))))))
            case "foo bar":
                _ = len(v2)
            case "a":
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 uint64 = 18446744073709551615
                    _ = fi18
                    _ = v19
                }
            default:
                if (1 != 0) {
                    switch (fi17 | (fi17 & -927)) {
                    case 32768:
                        v20 := GenS1{F0: -2147483648, F1: false}
                        _ = v20
                    default:
                        v7 = (v7 << ((18446744073709551615) & 63))
                    }
                }
            }
            var v21 string = ("" + "foo bar")
            v22 := GenS2{F0: fi17, F1: v1, F2: 181}
            _ = fi17
            _ = v21
            _ = v22
        }
        v5 = -95
    } else {
        v23 := GenS1{F0: 872, F1: (!(true && false))}
        v3 = "hello"
        _ = v23
    }
    var v24 float64 = v4
    v3 = ("" + (v3 + "a"))
    if (!false) {
        fmt.Println(v5, v24, v1)
        for fi25 := 0; fi25 < 5; fi25++ {
            if (!(1 < 18446744073709551615)) {
                fmt.Println(v7, v24, v1)
                switch v3 {
                case "a":
                    var v26 int = 42
                    _ = v26
                }
            }
            _ = fi25
        }
    } else {
        if true {
            _ = len(v2)
            v27 := GenS3{F0: 1, F1: (v3 + (("a" + "hello") + v3))}
            v28 := GenS3{F0: 32768, F1: "\n\t"}
            _ = v27
            _ = v28
        } else {
            for fi29 := 0; fi29 < 2; fi29++ {
                v16 -= (((v16 + float64(float32(v24))) + 0.0) * (-2.25 * ((v4 * 1.5) + 1e10)))
                _ = fi29
            }
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 uint8 = 1
                if false {
                    var v32 string = ((v3 + v3) + "hello")
                    _ = v32
                } else {
                    v16 -= (1.5 + (((v24 / (v24 + (v24 + v4))) / 3.14159) - v16))
                    v33 := []int{9223372036854775807, fi30, (fi30 >> ((97) & 63)), (fi30 & fi30)}
                    _ = v33
                }
                fmt.Println(v7, v3, v16)
                _ = fi30
                _ = v31
            }
        }
    }
    if ((v4 * v16) == (((v24 / v4) - v4) - float64(1.5))) {
        for fi34 := 0; fi34 < 4; fi34++ {
            v4 += 1e10
            fmt.Println(v7, v4, v24)
            switch fi34 {
            case 0:
                fmt.Println(v16, fi34, v24)
            case 2147483647:
                v35 := GenS4{F0: fi34, F1: -1}
                _ = v35
            }
            _ = fi34
        }
        for fi36 := 0; fi36 < 18; fi36++ {
            for fi37 := 0; fi37 < 19; fi37++ {
                v38 := []int{}
                var v39 int = (^(fi36))
                _ = fi37
                _ = v38
                _ = v39
            }
            if true {
                var v40 int = -3
                _ = v40
            }
            _ = fi36
        }
    }
    switch 127 {
    case 32767:
        v3 = ""
    }
    fmt.Println(v4, v1, v1)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
