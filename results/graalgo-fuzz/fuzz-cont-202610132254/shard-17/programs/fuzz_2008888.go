package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 int) int16 {
    p0 = 42
    switch (p0 % ((p0) | 1)) {
    case 2147483647:
        p0 = -2147483648
    case 128:
        var v6 float32 = ((1.5 + (float32(((1.5 / (float64(3.14159) - 0.0)) - (-2.25 / -2.25))) / (1e10 - (1.5 - (0.0 + 3.14159))))) - 1e10)
        _ = v6
    default:
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint = 18446744073709551615
            v9 := []int{fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v10 := GenS1{F0: 103, F1: 0}
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := []int{}
        v13 := GenS1{F0: 42, F1: -1}
        var v14 int32 = -250
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v10
    return -1
}

func main() {
    var v1 int32 = -1
    var v2 int64 = -153
    var v3 int32 = v1
    var v4 int = 0
    var v5 uint64 = 2
    v15 := GenS2{F0: -120, F1: 11}
    fmt.Println(v5, v3)
    v16 := []int{0, 32767, v4, -1}
    switch int(v2) {
    case 2147483647:
        v2 = v2
    }
    if false {
        if true {
            v5 = 797
            for fi17 := 0; fi17 < 2; fi17++ {
                for fi18 := 0; fi18 < 2; fi18++ {
                    switch (fi17 + (v4 * fi18)) {
                    case 2147483647:
                        _ = len(v16)
                    case 350:
                        var v19 uint16 = 781
                        _ = v19
                    case 42:
                        if len(v16) > 0 {
                            v16[0] = (fi17 / (((v4 | (fi17 + (fi17 >> ((42) & 63))))) | 1))
                        }
                    }
                    var v20 uint = 1
                    var v21 int16 = -1
                    _ = fi18
                    _ = v20
                    _ = v21
                }
                if true {
                    if false {
                        for fi22 := 0; fi22 < 19; fi22++ {
                            v16 = append(v16, (v4 - v4))
                            _ = fi22
                        }
                        fmt.Println(v4)
                    }
                } else {
                    v1 += (v3 << ((1) & 63))
                }
                _ = len(v16)
                _ = fi17
            }
            switch ("" + "hello") {
            case "":
                fmt.Println(v3, v5, v2)
            case "foo bar":
                var v23 uint = 0
                _ = v23
            case "\n\t":
                var v24 float64 = 0.1
                _ = v24
            }
        }
    }
    v25 := GenS1{F0: 42, F1: 9223372036854775807}
    if (-1 == v3) {
        var v26 uint8 = 128
        var v27 int = -1
        _ = v26
        _ = v27
    } else {
        v2 -= (v2 - (v2 + v2))
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        if false {
            v28 := GenS1{F0: 802, F1: (-(v2))}
            switch (v4 - v4) {
            case 879:
                for fi29 := 0; fi29 < 13; fi29++ {
                    v16 = append(v16, int(v2))
                    var v30 float32 = 3.14159
                    _ = fi29
                    _ = v30
                }
            }
            switch "a" {
            case "foo bar":
                v31 := GenS2{F0: 9, F1: 42}
                _ = v31
            case "hello":
                if len(v16) > 0 {
                    v16[0] = 4294967295
                }
            }
            _ = v28
        }
    }
    if false {
        if len(v16) > 0 {
            v16[0] = (v4 ^ (v4 >> ((307) & 63)))
        }
        v3 = (v1 % ((749) | 1))
    }
    v3 = (-(v3))
    v16 = append(v16, (v4 & (v4 | v4)))
    if (686 == 2147483648) {
        v5 -= 993
        v32 := GenS1{F0: 2, F1: (^(v2))}
        _ = v32
    }
    var v33 uint32 = 123
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 string = (("" + "hello") + "hello")
        var v36 int16 = -1
        _ = fi34
        _ = v35
        _ = v36
    }
    v37 := GenS1{F0: 0, F1: (-(v2))}
    for fi38 := 0; fi38 < 5; fi38++ {
        var v39 int64 = v2
        _ = fi38
        _ = v39
    }
    var v40 float64 = (3.14159 * float64(3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
