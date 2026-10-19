package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
func genFn0() string {
    v4 := GenS1{F0: (!true)}
    v5 := []int{127}
    var v6 float32 = ((-2.25 / ((1e10 - 1e10) - 0.1)) / ((float32(1e10) * -2.25) / 1e10))
    v7 := GenS1{F0: true}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{"foo bar", ((("a" + "") + "\n\t") + "a"), (("foo bar" + "\n\t") + "\n\t"), "foo bar"}
    var v3 string = ""
    if len(v2) > 0 {
        v2[0] = (v3 + v3)
    }
    var v8 int16 = -32768
    switch v3 {
    case "hello":
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 int32 = 2147483647
                    fmt.Println(v1, v11)
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
            v8 += v8
        }
    case "":
        v3 = ""
    case "\n\t":
        var v12 uint16 = 1
        _ = v12
    }
    var v13 string = ""
    switch -2147483648 {
    case 4294967295:
        if ((v8 << ((390) & 63)) != v8) {
            var v14 uint16 = 0
            var v15 float32 = 1.5
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v14
            _ = v15
        } else {
            if (69 <= 1) {
                fmt.Println(v8)
                v16 := GenS1{F0: false}
                _ = v16
            }
        }
    }
    v17 := GenS2{F0: 42}
    fmt.Println(v1)
    v3 = v3
    var v18 string = v13
    switch ((("\n\t" + v3) + v13) + v13) {
    case "":
        var v19 int16 = v8
        _ = v19
    case "a":
        var v20 uint32 = 2
        _ = v20
    case "foo bar":
        v3 = v13
    default:
        var v21 int32 = -2147483648
        _ = v21
    }
    _ = len(v2)
    switch "" {
    case "foo bar":
        v2 = append(v2, v18)
    case "":
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := GenS2{F0: 70}
            var v24 string = "\n\t"
            v24 = v18
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 uint32 = 284
        _ = fi25
        _ = v26
    }
    v1 -= 1e10
    v27 := GenS2{F0: 242}
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 int32 = 131
        var v30 bool = false
        switch v18 {
        case "":
            var v31 int = 32767
            _ = v31
        case "a":
            v3 = v18
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    var v32 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v8
}
