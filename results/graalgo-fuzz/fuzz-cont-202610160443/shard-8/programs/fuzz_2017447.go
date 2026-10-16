package main

import (
    "fmt"
)

func genFn0() string {
    var v6 float64 = 0.0
    var v7 int = -771
    fmt.Println(v6)
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    _ = v6
    _ = v7
    return ("\n\t" + ("hello" + ""))
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"foo bar", ("\n\t" + "a"), "", (("\n\t" + ("foo bar" + "foo bar")) + "\n\t")}
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := []int{0, 32767, 580, 365}
    var v9 bool = (0.0 > 1e10)
    v4["\n\t"] = 209
    v9 = (((("foo bar" + "a") + "") + (((("\n\t" + ("\n\t" + ("" + "hello"))) + (("a" + "\n\t") + ("a" + (("" + ("\n\t" + ("hello" + "foo bar"))) + "")))) + (("hello" + (((("a" + ("\n\t" + "\n\t")) + (("foo bar" + "hello") + "foo bar")) + "hello") + "a")) + (("foo bar" + "\n\t") + "a"))) + ("foo bar" + "\n\t"))) < "\n\t")
    for fi10 := 0; fi10 < 10; fi10++ {
        v9 = v9
        v9 = false
        _ = fi10
    }
    if (!false) {
        if v9 {
            var v11 int64 = 1
            _ = v11
        }
        var v12 float64 = 1.5
        _ = v12
    }
    if (42 > 526) {
        if len(v5) > 0 {
            v5[0] = 503
        }
        v9 = (-32768 > 0)
        switch ("hello" + "hello") {
        case "a":
            fmt.Println(v9)
        case "":
            fmt.Println(v9)
        case "hello":
            v9 = v9
        }
    }
    var v13 int8 = 127
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v9 = true
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
