package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 int = 315
    for fi4 := 0; fi4 < 9; fi4++ {
        v3 += fi4
        fmt.Println(v3, v3)
        v5 := GenS1{F0: 42, F1: "foo bar"}
        _ = fi4
        _ = v5
    }
    v3 = (v3 * v3)
    if (65535 > 665) {
        v2[(v3 | v3)] = ((("" + ("\n\t" + ("\n\t" + "a"))) + ("a" + "")) + "hello")
    } else {
        v6 := GenS1{F0: 65535, F1: "hello"}
        _ = v6
    }
    v3 += 730
    if ((!(!true)) && false) {
        for fi7 := 0; fi7 < 14; fi7++ {
            if false {
                v8 := GenS1{F0: 255, F1: "foo bar"}
                v9 := []int{v3, (fi7 / (((^(v3))) | 1))}
                v10 := GenS1{F0: 32767, F1: "a"}
                _ = v10
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        v11 := []int{0, 2147483647, 1}
        _ = v11
    } else {
        if false {
            v12 := GenS1{F0: 42, F1: ("\n\t" + "\n\t")}
            _ = v12
        }
        var v13 int = v3
        _ = v13
    }
    v14 := GenS1{F0: 65535, F1: ("\n\t" + ("\n\t" + ("a" + "")))}
    var v15 int64 = 9223372036854775807
    fmt.Println(v3, v15)
    v16 := []int{1, 9223372036854775807, (^(v3))}
    v17 := GenS2{F0: (v15 - v15), F1: 842, F2: (^(v15))}
    fmt.Println(v15, v15)
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS1{F0: 1, F1: "foo bar"}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v21 := GenS1{F0: 32767, F1: ("foo bar" + "hello")}
    for fi22 := 0; fi22 < 19; fi22++ {
        switch "" {
        case "foo bar":
            v23 := GenS2{F0: (-(v15)), F1: 442, F2: 9223372036854775807}
            _ = v23
        }
        if (0.0 >= ((1e10 + 1e10) + float64(float32(3.14159)))) {
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 bool = (!(float32(1.5) > 0.0))
                var v26 bool = v25
                _ = len(v2)
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        _ = fi22
    }
    v3 = (v3 * 127)
    fmt.Println(v15, v15)
    var v27 int8 = 57
    v16 = append(v16, 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v27
    _ = v3
}
