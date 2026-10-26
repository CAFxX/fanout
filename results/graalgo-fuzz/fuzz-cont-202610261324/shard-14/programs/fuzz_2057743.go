package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int16; F2 uint }

func main() {
    var v1 uint32 = 42
    v2 := []string{("foo bar" + "\n\t"), ((("hello" + "hello") + ("foo bar" + ("foo bar" + ""))) + ""), (((("foo bar" + "hello") + ("foo bar" + ((("foo bar" + "\n\t") + ("foo bar" + "\n\t")) + "hello"))) + "") + ("a" + "hello")), "hello"}
    var v3 bool = (127 >= 255)
    var v4 string = "foo bar"
    var v5 uint32 = 2
    v5 = (v5 << ((261) & 63))
    switch v4 {
    case "":
        v2 = append(v2, ("foo bar" + "\n\t"))
    case "a":
        if (4294967295 <= 18446744073709551615) {
            fmt.Println(v5)
        }
    default:
        v4 = v4
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        v3 = v3
        if true {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        } else {
            if ((!false) || (127 < 1)) {
                if (!v3) {
                    switch (fi6 - fi6) {
                    case 535:
                        var v7 uint = 0
                        _ = v7
                    default:
                        v8 := GenS1{F0: 255, F1: 42, F2: 0}
                        _ = v8
                    }
                    v9 := GenS2{F0: (v4 + v4)}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v9
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch v4 {
                case "foo bar":
                    for fi10 := 0; fi10 < 10; fi10++ {
                        if len(v2) > 0 {
                            v2[0] = v4
                        }
                        _ = fi10
                    }
                }
            } else {
                var v11 uint32 = 2147483647
                _ = v11
            }
        }
        _ = fi6
    }
    v12 := GenS3{F0: v5, F1: -314, F2: 1}
    var v13 int8 = -128
    var v14 uint64 = 322
    fmt.Println(v14, v4, v5)
    v3 = (!v3)
    fmt.Println(v13, v3, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
