package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint; F2 int8 }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{856}
    v3 := make(map[int]bool)
    var v4 float32 = 3.14159
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 0
    if false {
        if len(v2) > 0 {
            v2[0] = 0
        }
    }
    v4 = 0.1
    v2 = append(v2, 2147483647)
    v6 = 993
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 float32 = 1e10
        switch (("\n\t" + "\n\t") + "\n\t") {
        case "hello":
            for fi9 := 0; fi9 < 6; fi9++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v10 uint8 = 127
                var v11 uint32 = 121
                fmt.Println(v4, v10, v10)
                fmt.Println(v11)
                _ = fi9
                _ = v10
                _ = v11
            }
        case "foo bar":
            var v12 uint64 = 452
            _ = v12
        case "\n\t":
            switch ("a" + ("" + "\n\t")) {
            case "a":
                var v13 int16 = 1
                _ = v13
            case "":
                switch (("hello" + "hello") + "\n\t") {
                case "hello":
                    for fi14 := 0; fi14 < 17; fi14++ {
                        fmt.Println(fi14, v6, v4)
                        var v15 int8 = 0
                        var v16 int16 = 42
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
            }
        }
        _ = fi7
        _ = v8
    }
    v17 := GenS1{F0: 429, F1: (v6 | (v6 >> ((1) & 63))), F2: 1}
    v4 = v4
    _ = len(v2)
    v18 := v3[-497]
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v18)
        _ = fi19
    }
    var v20 uint32 = 2
    var v21 uint16 = (v6 & (v6 * 42))
    v20 += v20
    if (971 >= v21) {
        if v18 {
            v22 := GenS1{F0: 18446744073709551615, F1: v6, F2: -1}
            delete(v5, 42)
            _ = v22
        } else {
            if true {
                v21 -= v6
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 int16 = -637
                    v21 = v21
                    _ = fi23
                    _ = v24
                }
                _ = len(v5)
            } else {
                var v25 uint8 = 193
                _ = len(v1)
                v26 := GenS2{F0: (^(v20)), F1: 2, F2: 100}
                _ = v25
                _ = v26
            }
        }
    }
    delete(v1, ("\n\t" + "hello"))
    for fi27 := 0; fi27 < 17; fi27++ {
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := v5[v20]
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    v30 := GenS3{F0: ((v18 && (65535 >= 0)) && false)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
