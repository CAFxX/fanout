package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []string{((("\n\t" + ("foo bar" + ("foo bar" + ("\n\t" + "hello")))) + ("hello" + "")) + "\n\t"), "", ("\n\t" + ""), "hello"}
    v2 := make(map[string]int)
    var v3 float64 = (-2.25 * float64(0.0))
    switch -447 {
    case 4294967295:
        delete(v2, ("a" + (("\n\t" + (("hello" + "") + "a")) + "hello")))
    case 32768:
        if ((307 > 120) || (((2 >= 0) && ((false || ((103 != 127) && false)) || true)) && (100 < 42))) {
            for fi4 := 0; fi4 < 20; fi4++ {
                v3 = v3
                _ = fi4
            }
            v5 := GenS1{F0: 1000}
            v3 -= 0.1
            _ = v5
        }
    }
    v6 := GenS1{F0: 1000}
    var v7 int8 = 100
    v7 = v7
    var v8 uint8 = 128
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := v2[("foo bar" + "a")]
        _ = fi9
        _ = v10
    }
    switch "a" {
    case "hello":
        if (!false) {
            for fi11 := 0; fi11 < 6; fi11++ {
                v8 += v8
                _ = fi11
            }
        }
    }
    if true {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 int16 = 32767
            _ = len(v1)
            _ = fi12
            _ = v13
        }
    } else {
        v14 := GenS1{F0: -1}
        if true {
            var v15 uint = 196
            _ = v15
        }
        _ = v14
    }
    v8 = v8
    var v16 string = (("hello" + "foo bar") + "")
    v7 = (^(v7))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
