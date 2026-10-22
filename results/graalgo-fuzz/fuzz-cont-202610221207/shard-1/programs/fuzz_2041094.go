package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint8) uint8 {
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 2; fi5++ {
            p0 = (p0 + (p0 | 2))
            v6 := GenS1{F0: -150, F1: p0}
            p0 = p0
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    var v7 int = 0
    v8 := []int{}
    var v9 uint8 = (p0 ^ (p0 ^ p0))
    fmt.Println(p0, v7, v9)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 int8, p1 uint16) string {
    var v10 float64 = float64(0.0)
    fmt.Println(v10, p1)
    if false {
        fmt.Println(v10)
    }
    v10 += 3.14159
    v10 += (v10 - (((v10 + v10) * 0.1) / ((1.5 * v10) + 1e10)))
    _ = p0
    _ = p1
    _ = v10
    return "hello"
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 356
    v3 := make(map[string]uint32)
    if true {
        delete(v1, ("a" + ""))
        v11 := v1["\n\t"]
        if false {
            for fi12 := 0; fi12 < 1; fi12++ {
                switch ("foo bar" + "\n\t") {
                case "a":
                    v13 := []int{}
                    _ = v13
                case "":
                    fmt.Println(fi12)
                case "foo bar":
                    v14 := []int{}
                    _ = v14
                }
                var v15 bool = ((false && (3.14159 <= 3.14159)) || (false || false))
                v16 := GenS2{F0: v2, F1: 757, F2: ""}
                _ = fi12
                _ = v15
                _ = v16
            }
            fmt.Println(v11, v11)
        } else {
            v17 := GenS2{F0: (v2 - (v2 % (((^(v2))) | 1))), F1: 907, F2: ("hello" + "foo bar")}
            v2 -= (v2 / ((65535) | 1))
            _ = v17
        }
        _ = v11
    } else {
        v3["a"] = 2147483647
        v2 += (v2 >> ((441) & 63))
        var v18 string = "a"
        _ = v18
    }
    if false {
        switch 0 {
        case 32768:
            if true {
                v19 := GenS3{F0: 42}
                if (0 <= -2147483648) {
                    if (true && (!(true && (0 > 2147483647)))) {
                        v20 := GenS3{F0: 42}
                        v21 := GenS3{F0: 42}
                        v2 = 490
                        _ = v20
                        _ = v21
                    }
                }
                _ = v19
            }
        case -1:
            switch -9223372036854775808 {
            case 283:
                v22 := GenS4{F0: 42, F1: "hello"}
                _ = v22
            case 32768:
                v2 -= v2
            }
        }
        v23 := GenS4{F0: 518, F1: "hello"}
        v2 += v2
        _ = v23
    }
    v2 += (v2 - 255)
    v24 := GenS4{F0: 42, F1: "\n\t"}
    if false {
        v25 := []int{393}
        v2 = v2
        var v26 bool = (true && true)
        _ = v25
        _ = v26
    }
    v27 := GenS5{F0: -428}
    var v28 uint32 = 798
    v28 = (v28 << ((65535) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
}
