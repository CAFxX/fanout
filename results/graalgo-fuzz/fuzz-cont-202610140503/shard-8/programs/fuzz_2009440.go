package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    var v3 bool = (3.14159 != -2.25)
    var v4 uint = 839
    var v5 string = "\n\t"
    v3 = v3
    v4 = v4
    fmt.Println(v4, v3, v5)
    fmt.Println(v5, v5, v3)
    switch ("hello" + "a") {
    case "\n\t":
        fmt.Println(v4, v4, v3)
    default:
        _ = len(v1)
    }
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint16 = 548
        var v9 string = ((v5 + "foo bar") + "\n\t")
        v3 = false
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 433
    v4 += 37
    if v3 {
        fmt.Println(v3)
    }
    var v11 float32 = 0.0
    for fi12 := 0; fi12 < 2; fi12++ {
        switch fi12 {
        case 380:
            var v13 int8 = 43
            _ = v13
        case -481:
            fmt.Println(v5, fi12)
        case -916:
            switch "" {
            case "":
                if false {
                    _ = len(v2)
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 string = v5
                        if (42 < 139) {
                            _ = len(v1)
                            v11 -= v6
                        }
                        _ = fi14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 13; fi16++ {
                        if (v3 || true) {
                            v1 = append(v1, v5)
                        } else {
                            fmt.Println(fi16, v6, fi12)
                            var v17 int64 = 659
                            _ = v17
                        }
                        _ = fi16
                    }
                } else {
                    var v18 uint16 = 0
                    v19 := GenS1{F0: v5}
                    _ = v18
                    _ = v19
                }
            case "\n\t":
                var v20 uint16 = 1
                _ = v20
            }
        default:
            var v21 int = (fi12 - (-(fi12)))
            _ = v21
        }
        _ = fi12
    }
    v5 = (v5 + (("\n\t" + v5) + (v5 + ((("" + (v5 + "a")) + v5) + v5))))
    v5 = "hello"
    delete(v2, -640)
    v6 += v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
