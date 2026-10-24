package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint32 = 852
    var v2 string = "hello"
    var v3 bool = false
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 255, F1: 0}
    fmt.Println(v1, v3, v1)
    fmt.Println(v1, v2, v2)
    var v6 uint32 = 0
    fmt.Println(v6)
    v7 := GenS1{F0: 2, F1: -2147483648}
    delete(v4, (v6 % ((2147483648) | 1)))
    v6 -= 2
    fmt.Println(v2)
    if v3 {
        if (537 != 0) {
            v1 -= v1
        } else {
            var v8 float64 = 3.14159
            _ = v8
        }
        var v9 int16 = 0
        _ = v9
    }
    v4[(v6 % (((v6 * 0)) | 1))] = 32768
    _ = len(v4)
    switch ("foo bar" + ("a" + "\n\t")) {
    case "a":
        v2 = v2
    default:
        if (!v3) {
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint = 2
                var v12 uint16 = 2
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 float32 = (1e10 + ((3.14159 / (float32(1e10) - 0.1)) / float32(3.14159)))
            _ = v13
        } else {
            if (!(-9223372036854775808 == -732)) {
                var v14 int8 = 42
                v2 = ((v2 + v2) + "")
                if (0 >= 0) {
                    v3 = (!false)
                    if (1.5 == (float64(0.1) * -2.25)) {
                        switch ((((v2 + "hello") + (v2 + "foo bar")) + v2) + "hello") {
                        case "a":
                            v15 := []int{32768}
                            _ = v15
                        case "foo bar":
                            fmt.Println(v1)
                        }
                        for fi16 := 0; fi16 < 4; fi16++ {
                            fmt.Println(fi16, v14)
                            fmt.Println(v14)
                            _ = fi16
                        }
                        v1 += v6
                    }
                    for fi17 := 0; fi17 < 16; fi17++ {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            var v19 float32 = ((float32((0.1 - 3.14159)) * (1.5 - 0.1)) / (3.14159 * (0.0 + (-2.25 - 0.1))))
                            var v20 int64 = 9223372036854775807
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        _ = fi17
                    }
                } else {
                    v14 = v14
                }
                _ = v14
            }
            v1 += (v6 * 65535)
        }
    }
    v1 = (v1 % (((v1 - (^(v1)))) | 1))
    v1 -= 0
    v21 := []int{}
    switch v2 {
    case "foo bar":
        for fi22 := 0; fi22 < 8; fi22++ {
            v6 -= v1
            switch (^(fi22)) {
            case 2147483647:
                v1 = v1
            }
            _ = fi22
        }
    case "a":
        var v23 uint16 = 65535
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
