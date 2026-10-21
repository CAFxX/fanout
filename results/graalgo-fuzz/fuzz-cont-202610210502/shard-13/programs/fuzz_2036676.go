package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 string; F1 int64; F2 int8 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []string{"hello", ("\n\t" + ("" + (("hello" + "a") + ""))), (("foo bar" + ("a" + ("\n\t" + (("\n\t" + "hello") + "a")))) + ("hello" + ("hello" + "hello"))), ""}
    v2 := []uint32{378, 218, 1, 2}
    var v3 int8 = 127
    var v4 int64 = 9223372036854775807
    var v5 int32 = -2147483648
    var v6 uint = 940
    var v7 int32 = 1
    var v8 int8 = (v3 & v3)
    v7 += v5
    v6 = v6
    v5 += v5
    _ = len(v1)
    for fi9 := 0; fi9 < 7; fi9++ {
        if true {
            switch (fi9 + fi9) {
            case 2147483647:
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 int16 = 32767
                    var v12 uint = v6
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case -2147483648:
                v13 := GenS1{F0: 0, F1: 0}
                _ = v13
            }
            var v14 uint16 = 0
            _ = v14
        } else {
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            for fi15 := 0; fi15 < 13; fi15++ {
                if (false || (v6 <= (v6 * (v6 % (((v6 | v6)) | 1))))) {
                    if (2 == (v6 ^ (v6 - v6))) {
                        switch (fi9 * (^(fi15))) {
                        case -2147483648:
                            v6 -= v6
                            v16 := GenS2{F0: "hello", F1: 1, F2: v8}
                            _ = v16
                        }
                        fmt.Println(v4, v3, v8)
                    }
                }
                if (false && (42 > -32768)) {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    switch "a" {
                    case "\n\t":
                        v2 = append(v2, 299)
                    }
                    v17 := GenS3{F0: v4}
                    _ = v17
                }
                _ = fi15
            }
            if ((3.14159 - 0.0) < ((0.1 * 0.0) / 0.1)) {
                var v18 int64 = v4
                _ = v18
            } else {
                switch ("a" + ("a" + "")) {
                case "foo bar":
                    fmt.Println(fi9, v7)
                case "":
                    v6 -= (^(v6))
                case "\n\t":
                    for fi19 := 0; fi19 < 2; fi19++ {
                        v1 = append(v1, ("\n\t" + ("" + ("" + "a"))))
                        _ = fi19
                    }
                }
                for fi20 := 0; fi20 < 18; fi20++ {
                    var v21 int8 = v3
                    v8 -= (v3 | 123)
                    _ = fi20
                    _ = v21
                }
                var v22 int64 = v4
                _ = v22
            }
        }
        _ = fi9
    }
    v4 += v4
    var v23 float64 = float64(-2.25)
    v5 = v7
    var v24 float32 = (3.14159 / 1.0)
    v7 -= v5
    _ = len(v2)
    switch "hello" {
    case "\n\t":
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
