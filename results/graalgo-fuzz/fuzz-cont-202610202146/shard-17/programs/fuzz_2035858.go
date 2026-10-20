package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := []string{("a" + "\n\t"), "foo bar"}
    var v2 bool = ((1e10 + 1e10) <= (3.14159 * float32(float64(float32(0.0)))))
    var v3 uint64 = 4294967295
    v4 := GenS1{F0: 42}
    v3 = v3
    for fi5 := 0; fi5 < 17; fi5++ {
        if (379 < 0) {
            if (606 != -32768) {
                if (!(-2147483648 != -1)) {
                    _ = len(v1)
                    v1 = append(v1, "")
                    if (((42 < 1) || (169 == 1)) || (((0.0 - (1.5 * float32(1.5))) * 1e10) > ((float32((0.1 + float64(float32(-2.25)))) - ((float32(0.0) * -2.25) / 3.14159)) + 0.0))) {
                        v3 += 2
                        v2 = v2
                    }
                }
            } else {
                _ = len(v1)
                switch (("foo bar" + "foo bar") + "a") {
                case "foo bar":
                    var v6 string = "hello"
                    _ = v6
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            var v7 bool = v2
            var v8 int32 = 42
            _ = v7
            _ = v8
        } else {
            switch "\n\t" {
            case "a":
                v9 := GenS2{F0: -11, F1: 329}
                _ = v9
            case "foo bar":
                var v10 float64 = (0.1 - 3.14159)
                _ = v10
            default:
                var v11 uint8 = 214
                _ = v11
            }
            switch (fi5 ^ fi5) {
            case 1:
                var v12 float64 = 0.0
                _ = v12
            case 32767:
                var v13 bool = (241 <= 2147483647)
                _ = v13
            default:
                v3 = 4294967295
            }
            v14 := GenS3{F0: 0, F1: v2}
            _ = v14
        }
        var v15 uint = 65535
        v2 = (!v2)
        _ = fi5
        _ = v15
    }
    v2 = (4294967295 < -2147483648)
    v2 = v2
    fmt.Println(v3, v3)
    v3 = v3
    v16 := GenS3{F0: 42, F1: true}
    switch "\n\t" {
    case "hello":
        var v17 int64 = -800
        _ = v17
    }
    fmt.Println(v3, v2)
    if v2 {
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS3{F0: (v3 * (^(v3))), F1: (42 <= 128)}
            _ = fi18
            _ = v19
        }
        v20 := GenS4{F0: v3}
        for fi21 := 0; fi21 < 16; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                v23 := GenS1{F0: 32767}
                _ = fi22
                _ = v23
            }
            _ = len(v1)
            for fi24 := 0; fi24 < 3; fi24++ {
                switch fi21 {
                case 73:
                    var v25 uint64 = v3
                    _ = v25
                }
                _ = fi24
            }
            _ = fi21
        }
        _ = v20
    } else {
        v2 = false
    }
    var v26 uint16 = 920
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
