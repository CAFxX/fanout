package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() bool {
    var v4 uint64 = 42
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(v4, v4)
        _ = fi5
    }
    var v6 uint8 = 128
    switch 1 {
    case -989:
        v4 = v4
    default:
        if (53 > 42) {
            v4 = 42
            v6 = 80
        } else {
            switch ((((("a" + "a") + "hello") + ((("foo bar" + "hello") + "") + "a")) + ("" + ("a" + ("foo bar" + ("\n\t" + ("a" + "")))))) + ("foo bar" + "a")) {
            case "":
                var v7 string = ((("\n\t" + "") + (((((("\n\t" + "a") + "") + ("" + ("" + ("foo bar" + "")))) + "hello") + "hello") + "hello")) + (("a" + "hello") + "\n\t"))
                _ = v7
            default:
                v6 = v6
            }
            v8 := GenS1{F0: 127}
            var v9 int32 = -381
            _ = v8
            _ = v9
        }
    }
    _ = v4
    _ = v6
    return (0 <= 545)
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ("foo bar" + ("hello" + ("" + "\n\t")))
    v3 := make(map[uint32]string)
    switch v2 {
    case "hello":
        for fi10 := 0; fi10 < 20; fi10++ {
            switch (fi10 * 127) {
            case -361:
                v2 = (v2 + v2)
            }
            fmt.Println(fi10)
            v11 := []int{fi10, 1, (fi10 << ((215) & 63))}
            _ = fi10
            _ = v11
        }
    case "":
        var v12 string = ""
        _ = v12
    default:
        v13 := []int{230}
        _ = v13
    }
    v14 := []int{448, 32768, -712, 9223372036854775807}
    fmt.Println(v2)
    v15 := v3[0]
    if (283 != 42) {
        v3[344] = v15
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint32 = 370
            _ = fi16
            _ = v17
        }
    }
    v15 = v2
    v18 := GenS1{F0: 127}
    v15 = v15
    v14 = append(v14, -963)
    var v19 uint32 = 2
    var v20 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
