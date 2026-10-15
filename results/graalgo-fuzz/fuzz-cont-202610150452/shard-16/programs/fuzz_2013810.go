package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []uint32{459}
    v2 := make(map[string]string)
    var v3 int32 = 1
    var v4 uint64 = 4294967295
    fmt.Println(v4, v3)
    var v5 uint = 255
    var v6 int64 = -79
    delete(v2, "a")
    if (true || false) {
        v7 := GenS1{F0: "a"}
        v5 = (v5 | uint(v4))
        var v8 uint16 = 90
        _ = v7
        _ = v8
    } else {
        fmt.Println(v4)
        if ((("a" + ((("hello" + ("a" + "a")) + (("\n\t" + (("foo bar" + "a") + ((("\n\t" + "hello") + (("a" + "hello") + (("hello" + ("\n\t" + (("\n\t" + "\n\t") + ""))) + "hello"))) + "hello"))) + "")) + "a")) + "hello") >= "foo bar") {
            var v9 int8 = 127
            v3 = v3
            if true {
                v1 = append(v1, 2)
                var v10 int32 = 0
                if true {
                    for fi11 := 0; fi11 < 4; fi11++ {
                        if false {
                            v5 += 42
                            v9 += v9
                            v1 = append(v1, 0)
                        } else {
                            if true {
                                v6 += v6
                                v3 = (v3 ^ (v3 % ((v3) | 1)))
                            } else {
                                var v12 uint16 = 1
                                _ = v12
                            }
                            v4 = 2
                        }
                        _ = len(v1)
                        _ = fi11
                    }
                }
                _ = v10
            }
            _ = v9
        } else {
            var v13 uint64 = v4
            v4 -= (v4 ^ 357)
            v5 = v5
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v3 = (^(v3))
        _ = fi14
    }
    switch 4294967295 {
    case -547:
        v15 := GenS1{F0: ("\n\t" + ((("hello" + "") + "\n\t") + "foo bar"))}
        _ = v15
    }
    var v16 int16 = 0
    for fi17 := 0; fi17 < 14; fi17++ {
        v5 = 625
        _ = fi17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
