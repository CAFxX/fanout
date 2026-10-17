package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    var v5 int = -719
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: ("hello" + ("\n\t" + ""))}
        v5 += 259
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 string = "a"
        v9 = v9
        var v10 int16 = 1000
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if (392 < 255) {
            delete(v4, 2147483647)
        }
        if (-2.25 > (float32(3.14159) + 0.1)) {
            fmt.Println(fi11, v5)
            v12 := v2[4294967295]
            v13 := []int{(fi11 & 32768), fi11, (v5 ^ 127)}
            _ = v12
            _ = v13
        } else {
            v14 := []int{}
            v5 += -884
            v1["\n\t"] = 42
            _ = v14
        }
        _ = fi11
    }
    v5 -= (-(v5))
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: "\n\t"}
        for fi17 := 0; fi17 < 12; fi17++ {
            if false {
                var v18 uint16 = 644
                _ = len(v2)
                var v19 int = fi15
                if (!true) {
                    if true {
                        v18 = (v18 + v18)
                    }
                    v20 := []int{v19}
                    v20 = append(v20, 1)
                    _ = v20
                } else {
                    v18 -= v18
                }
                _ = v18
                _ = v19
            } else {
                fmt.Println(fi17)
                v21 := GenS1{F0: ""}
                var v22 int32 = -365
                _ = v21
                _ = v22
            }
            fmt.Println(fi17, fi15)
            _ = fi17
        }
        _ = fi15
        _ = v16
    }
    v23 := GenS1{F0: (((("\n\t" + "") + "") + ("hello" + "foo bar")) + "hello")}
    fmt.Println(v5)
    _ = len(v2)
    for fi24 := 0; fi24 < 2; fi24++ {
        delete(v3, 194)
        switch "a" {
        case "hello":
            v5 += v5
        case "\n\t":
            var v25 int = 127
            _ = v25
        case "foo bar":
            v26 := []int{(v5 - (v5 >> ((255) & 63))), v5, 127, fi24}
            _ = v26
        default:
            fmt.Println(fi24, v5)
        }
        var v27 int16 = -1
        _ = fi24
        _ = v27
    }
    fmt.Println(v5)
    v28 := GenS1{F0: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
