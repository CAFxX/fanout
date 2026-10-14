package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int8, p1 int16) uint {
    p0 = (p0 & (p0 - (^(p0))))
    var v5 uint64 = 42
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -57
    v3 := []string{"", "", ((("\n\t" + "hello") + ("hello" + "a")) + (("a" + ((("\n\t" + "") + "\n\t") + "a")) + ""))}
    v4 := []uint32{809, 4294967295}
    v2 = v2
    v2 -= 87
    if (("\n\t" + "foo bar") > "") {
        if false {
            for fi6 := 0; fi6 < 14; fi6++ {
                v7 := GenS1{F0: 708}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi6
                _ = v7
            }
            switch (("\n\t" + "\n\t") + (("foo bar" + "a") + "hello")) {
            case "hello":
                for fi8 := 0; fi8 < 20; fi8++ {
                    v2 += 1
                    v2 = 42
                    _ = fi8
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        var v9 uint = 255
        _ = v9
    }
    v10 := GenS1{F0: 1}
    v2 = v2
    fmt.Println(v2)
    if true {
        v2 += v2
    }
    fmt.Println(v2)
    v2 += (^(v2))
    var v11 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
