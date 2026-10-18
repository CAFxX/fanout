package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint16 }

func main() {
    var v1 uint32 = 2
    var v2 int8 = 0
    v3 := []uint32{2147483648, v1}
    v4 := []string{"a", ("hello" + "\n\t"), ((("foo bar" + "") + "") + ("foo bar" + "a"))}
    v3 = append(v3, v1)
    fmt.Println(v1, v2)
    var v5 string = "foo bar"
    for fi6 := 0; fi6 < 1; fi6++ {
        if len(v3) > 0 {
            v3[0] = (v1 + v1)
        }
        var v7 float32 = 1.5
        _ = fi6
        _ = v7
    }
    if (((3.14159 / 1e10) > ((3.14159 - 3.14159) * (1e10 / (0.0 - (1e10 * 1e10))))) || (!((!(-1 == 1)) || false))) {
        var v8 uint8 = 146
        v5 = "hello"
        _ = v8
    }
    fmt.Println(v2, v2)
    switch 0 {
    case 68:
        switch "\n\t" {
        case "a":
            v5 = (v5 + (((v5 + v5) + v5) + "hello"))
        case "foo bar":
            v5 = ""
        case "hello":
            switch 2147483647 {
            case 9223372036854775807:
                var v9 uint8 = 255
                _ = v9
            case -425:
                fmt.Println(v5, v5)
            case -88:
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch (v5 + "") {
                    case "hello":
                        if (-1 != 42) {
                            var v11 int16 = -1
                            if false {
                                switch ((v5 + v5) + "\n\t") {
                                case "\n\t":
                                    for fi12 := 0; fi12 < 16; fi12++ {
                                        switch (fi12 / ((644) | 1)) {
                                        case 128:
                                            v13 := GenS1{F0: 42, F1: -2147483648}
                                            _ = v13
                                        case 561:
                                            v11 += 42
                                        default:
                                            var v14 int = fi10
                                            _ = v14
                                        }
                                        _ = fi12
                                    }
                                case "foo bar":
                                    fmt.Println(fi10, v5)
                                case "":
                                    if true {
                                        v1 = (v1 % ((v1) | 1))
                                        v15 := GenS2{F0: v11, F1: v5, F2: 42}
                                        v4 = append(v4, (("" + (((v5 + (("" + v5) + ((v5 + v5) + ("\n\t" + v5)))) + "hello") + (v5 + "\n\t"))) + v5))
                                        _ = v15
                                    }
                                }
                            }
                            v16 := GenS1{F0: 116, F1: 42}
                            _ = v11
                            _ = v16
                        }
                    case "\n\t":
                        var v17 int8 = v2
                        _ = v17
                    default:
                        var v18 string = v5
                        _ = v18
                    }
                    _ = fi10
                }
            }
        default:
            var v19 int = -1
            _ = v19
        }
    }
    v20 := GenS3{F0: 128, F1: (v1 + 910), F2: 32767}
    v21 := GenS3{F0: 231, F1: (v1 + (v1 / ((65535) | 1))), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
