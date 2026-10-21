package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []int{2147483647, -1, 982, -62}
    var v2 string = ("a" + "a")
    v3 := []string{"", (v2 + (("foo bar" + v2) + v2)), (((("" + (("\n\t" + (((("foo bar" + ("\n\t" + "")) + (v2 + ((v2 + "a") + ((v2 + (v2 + v2)) + (("\n\t" + "hello") + v2))))) + ((v2 + "\n\t") + v2)) + (v2 + v2))) + "a")) + ((((((v2 + v2) + v2) + "foo bar") + ((((v2 + "a") + (("foo bar" + (v2 + (v2 + v2))) + v2)) + v2) + v2)) + v2) + (v2 + v2))) + "hello") + ""), (v2 + "foo bar")}
    v4 := []uint32{2, 2147483647, 65535}
    v3 = append(v3, ((("foo bar" + "hello") + v2) + "a"))
    var v5 uint32 = 641
    if true {
        var v6 uint64 = 1
        v7 := GenS1{F0: ""}
        v5 -= 1
        _ = v6
        _ = v7
    }
    v2 = (((v2 + v2) + v2) + "\n\t")
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: v2}
        _ = fi8
        _ = v9
    }
    var v10 float32 = -2.25
    v5 = v5
    v10 -= v10
    for fi11 := 0; fi11 < 20; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v12 bool = false
        switch (fi11 | 128) {
        case -1:
            v13 := GenS1{F0: v2}
            _ = v13
        default:
            if len(v1) > 0 {
                v1[0] = 905
            }
        }
        _ = fi11
        _ = v12
    }
    _ = len(v4)
    v14 := GenS1{F0: (v2 + (("hello" + v2) + v2))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
