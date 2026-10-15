package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16) int32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            p0 = 2
            p0 = p0
            _ = fi7
        }
        switch "a" {
        case "a":
            v8 := []int{fi6, fi6, 4294967295}
            _ = v8
        case "":
            v9 := GenS1{F0: 42}
            _ = v9
        case "foo bar":
            if true {
                p0 += p0
            } else {
                p0 = p0
            }
        }
        _ = fi6
    }
    p0 += (p0 + (p0 << ((4294967295) & 63)))
    var v10 int64 = 681
    v10 = v10
    _ = p0
    _ = v10
    return -1
}

func main() {
    var v1 uint16 = 42
    v2 := []int{-323, -760}
    v3 := []string{}
    var v4 int32 = 0
    v5 := []int{0}
    fmt.Println(v1)
    var v11 int16 = 0
    fmt.Println(v4)
    var v12 int = 807
    var v13 string = (((("a" + "\n\t") + (("\n\t" + "a") + (("foo bar" + "") + (("a" + "foo bar") + ("a" + "hello"))))) + "a") + "\n\t")
    v2 = append(v2, v12)
    v13 = (((v13 + "\n\t") + ((v13 + "") + (v13 + "\n\t"))) + "a")
    if false {
        _ = len(v5)
        var v14 float32 = 3.14159
        v12 += -976
        _ = v14
    }
    _ = len(v2)
    if true {
        v15 := GenS1{F0: 0}
        _ = v15
    }
    var v16 uint = 26
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
