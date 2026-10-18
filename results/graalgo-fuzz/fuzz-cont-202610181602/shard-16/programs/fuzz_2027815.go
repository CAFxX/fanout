package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int8 = -1
    var v2 uint8 = 42
    v3 := make(map[string]uint32)
    var v4 bool = (!(!true))
    var v5 uint = 255
    v1 += v1
    if false {
        v6 := GenS1{F0: v4}
        _ = v6
    }
    v7 := []int{32768, 32768}
    v5 -= 0
    for fi8 := 0; fi8 < 11; fi8++ {
        switch "hello" {
        case "a":
            var v9 string = "\n\t"
            _ = v9
        case "hello":
            for fi10 := 0; fi10 < 17; fi10++ {
                fmt.Println(v5, v4)
                v2 = (v2 + uint8(v5))
                var v11 uint64 = uint64(v2)
                _ = fi10
                _ = v11
            }
        case "foo bar":
            v4 = true
        default:
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 float32 = (-2.25 + (0.1 * (float32(1.5) - ((((-2.25 * (-2.25 + 0.0)) - 0.1) * 1.5) - 1e10))))
                switch (("hello" + "") + "\n\t") {
                case "foo bar":
                    if v4 {
                        if (!(4294967295 != 643)) {
                            v2 -= v2
                        }
                        v13 -= (float32(0.1) - v13)
                        v4 = (true && (552 > 2))
                    }
                default:
                    var v14 int16 = 32767
                    _ = v14
                }
                if (((v5 << (((^(v5))) & 63)) == v5) && v4) {
                    switch 667 {
                    case 32768:
                        var v15 float64 = ((float64(0.0) * (0.0 - float64(3.14159))) - 0.1)
                        _ = v15
                    case 42:
                        v16 := GenS1{F0: v4}
                        _ = v16
                    }
                    switch "a" {
                    case "\n\t":
                        var v17 uint32 = 851
                        _ = v17
                    }
                    v18 := GenS1{F0: v4}
                    _ = v18
                } else {
                    v13 += (float32((0.0 - 3.14159)) + (v13 / 0.1))
                }
                switch fi12 {
                case 4294967295:
                    fmt.Println(fi12)
                default:
                    v5 += 2
                }
                _ = fi12
                _ = v13
            }
        }
        v19 := GenS1{F0: v4}
        for fi20 := 0; fi20 < 16; fi20++ {
            switch fi8 {
            case 42:
                _ = len(v7)
            case 32:
                var v21 uint8 = 255
                v22 := GenS1{F0: (!v4)}
                var v23 int16 = 0
                delete(v3, "")
                v5 += v5
                var v24 uint64 = 0
                fmt.Println(v5)
                _ = v21
                _ = v22
                _ = v23
                _ = v24
            case -932:
                v7 = append(v7, (fi20 ^ (fi20 & (fi8 % ((-832) | 1)))))
            default:
                v5 -= (v5 / (((v5 ^ (v5 + 65535))) | 1))
            }
            v4 = ((int32(v2) <= 622) || ("a" >= "a"))
            switch ("\n\t" + ("hello" + "foo bar")) {
            case "a":
                v1 += (-(v1))
            }
            _ = fi20
        }
        _ = fi8
        _ = v19
    }
    if v4 {
        v5 += 860
    } else {
        var v25 int = -2147483648
        v2 = (^(v2))
        switch "foo bar" {
        case "":
            v4 = v4
        default:
            for fi26 := 0; fi26 < 11; fi26++ {
                if v4 {
                    v2 = (^(v2))
                }
                _ = fi26
            }
        }
        _ = v25
    }
    fmt.Println(v1, v5)
    if true {
        v27 := GenS1{F0: (65535 >= 4294967295)}
        v5 = (v5 % ((v5) | 1))
        switch 42 {
        case -9223372036854775808:
            v7 = append(v7, 127)
        case 166:
            v3[""] = 4294967295
        }
        _ = v27
    } else {
        for fi28 := 0; fi28 < 4; fi28++ {
            v29 := v3["a"]
            if v4 {
                if true {
                    v4 = (!v4)
                    var v30 uint16 = 182
                    _ = v30
                }
            }
            _ = fi28
            _ = v29
        }
        var v31 bool = v4
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
