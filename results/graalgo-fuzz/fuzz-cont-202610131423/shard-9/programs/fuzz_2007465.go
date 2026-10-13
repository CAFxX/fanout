package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]int)
    v3 := []string{("" + ""), "\n\t", "hello", "a"}
    fmt.Println(v1)
    for fi4 := 0; fi4 < 18; fi4++ {
        if false {
            if (4294967295 < 18446744073709551615) {
                switch "a" {
                case "foo bar":
                    switch "" {
                    case "hello":
                        if false {
                            switch (fi4 * (fi4 | -9223372036854775808)) {
                            case 1:
                                var v5 float64 = (3.14159 + (0.1 - (1.5 - ((0.0 / float64(float32(1.5))) - 1e10))))
                                _ = v5
                            case -809:
                                var v6 uint = 52
                                _ = v6
                            default:
                                var v7 uint16 = 65535
                                _ = v7
                            }
                            _ = len(v2)
                            switch (fi4 | 1) {
                            case 9223372036854775807:
                                var v8 uint64 = 312
                                _ = v8
                            case 32768:
                                delete(v2, "\n\t")
                            }
                        }
                    case "a":
                        var v9 uint = 255
                        _ = v9
                    }
                case "":
                    var v10 int64 = 42
                    _ = v10
                case "a":
                    v11 := GenS1{F0: 9223372036854775807}
                    _ = v11
                default:
                    v1 = 1
                }
                v1 = (v1 * 4294967295)
            }
            var v12 string = "hello"
            _ = v12
        }
        _ = fi4
    }
    if (-2147483648 == 0) {
        if false {
            fmt.Println(v1)
            for fi13 := 0; fi13 < 8; fi13++ {
                v2["foo bar"] = 32767
                v1 = 232
                _ = fi13
            }
            switch 128 {
            case 32768:
                fmt.Println(v1)
            case -597:
                for fi14 := 0; fi14 < 4; fi14++ {
                    v15 := GenS1{F0: 752}
                    _ = fi14
                    _ = v15
                }
            case 127:
                v1 += v1
            default:
                fmt.Println(v1)
            }
        }
        v16 := GenS1{F0: 583}
        switch "foo bar" {
        case "a":
            if true {
                var v17 int32 = 1
                _ = v17
            }
        case "":
            fmt.Println(v1)
        }
        _ = v16
    }
    v1 -= 65535
    v1 = v1
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 bool = (!(true && false))
        _ = fi18
        _ = v19
    }
    v1 -= (v1 | (v1 >> ((0) & 63)))
    fmt.Println(v1)
    var v20 int = 128
    v20 += v20
    delete(v2, "hello")
    v20 += v20
    var v21 bool = ((100 < -1) && (!false))
    v21 = ("hello" == "a")
    fmt.Println(v21, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
