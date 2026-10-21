package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    var v1 int = 373
    var v2 int8 = 0
    v3 := make(map[string]int)
    v4 := make(map[int]uint32)
    v5 := []int{(v1 + 0), int(v2), (^(v1)), (v1 ^ (v1 - int(v2)))}
    _ = len(v3)
    delete(v3, ("hello" + "\n\t"))
    fmt.Println(v1, v1)
    var v6 uint32 = 0
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            var v8 string = ""
            _ = v8
        }
        if (809 < 1) {
            v1 -= (v1 << ((4294967295) & 63))
        } else {
            if true {
                var v9 bool = true
                if v9 {
                    switch (("a" + (("\n\t" + ("a" + (("hello" + "a") + ("" + ("" + "foo bar"))))) + ("\n\t" + "foo bar"))) + "") {
                    case "a":
                        for fi10 := 0; fi10 < 3; fi10++ {
                            _ = len(v5)
                            v9 = (255 != 163)
                            v1 = (fi7 & (-(v1)))
                            _ = fi10
                        }
                    }
                    delete(v4, v1)
                }
                _ = v9
            }
            if (false && ((!true) && (false || (!("hello" >= (("" + ("hello" + "\n\t")) + ("foo bar" + ""))))))) {
                v1 += v1
            } else {
                v11 := GenS1{F0: (("foo bar" + "a") + "hello"), F1: v2}
                v12 := GenS1{F0: "", F1: v2}
                _ = v11
                _ = v12
            }
        }
        _ = len(v5)
        _ = fi7
    }
    fmt.Println(v2, v1)
    var v13 float64 = (3.14159 * ((3.14159 / 1.0) * -2.25))
    v2 = v2
    v13 -= (v13 + v13)
    fmt.Println(v6, v1, v2)
    var v14 uint64 = 1
    fmt.Println(v1, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
