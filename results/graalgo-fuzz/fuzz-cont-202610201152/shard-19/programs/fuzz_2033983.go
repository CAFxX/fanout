package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int16, p1 int8, p2 int8) int {
    p0 = p0
    var v5 int16 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 128
}

func main() {
    var v1 int64 = -1
    v2 := []string{"foo bar", "", ("hello" + ("a" + "a"))}
    var v3 uint64 = 20
    var v4 uint = 174
    v6 := GenS1{F0: "", F1: -415}
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v9 int16 = 0
            _ = fi8
            _ = v9
        }
        if true {
            for fi10 := 0; fi10 < 17; fi10++ {
                switch (("foo bar" + "foo bar") + "a") {
                case "":
                    if (!true) {
                        v4 = 65535
                        var v11 float64 = (1e10 - 0.0)
                        _ = v11
                    } else {
                        v12 := GenS1{F0: "\n\t", F1: 1}
                        v2 = append(v2, "hello")
                        v3 = (v3 % (((v3 ^ v3)) | 1))
                        _ = v12
                    }
                case "\n\t":
                    v1 -= (v1 << (((v4 | (v4 / ((v4) | 1)))) & 63))
                default:
                    if false {
                        v2 = append(v2, ("\n\t" + ("\n\t" + (((("\n\t" + "a") + "a") + "a") + "hello"))))
                    }
                }
                _ = len(v2)
                _ = fi10
            }
            v13 := GenS1{F0: ("\n\t" + "hello"), F1: -1}
            v4 += (^(v4))
            _ = v13
        } else {
            if (100 <= -3) {
                switch "hello" {
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                case "a":
                    fmt.Println(v3)
                }
                var v14 float64 = 0.1
                _ = v14
            }
        }
        _ = fi7
    }
    var v15 int64 = (v1 ^ v1)
    _ = len(v2)
    v3 -= 606
    fmt.Println(v4, v4, v15)
    v2 = append(v2, "hello")
    v16 := GenS2{F0: 4294967295}
    v15 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
