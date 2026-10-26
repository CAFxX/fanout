package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint = 39
    var v2 bool = (1 < 42)
    var v3 uint64 = 2
    v4 := []string{"", (("a" + "foo bar") + "hello"), "\n\t", "foo bar"}
    v5 := []string{"hello", "\n\t", "a", "\n\t"}
    v6 := GenS1{F0: int16(v3)}
    v5 = append(v5, "\n\t")
    v2 = (false && (!(v2 && v2)))
    fmt.Println(v2, v1)
    var v7 float64 = -2.25
    fmt.Println(v1)
    fmt.Println(v7, v7)
    switch 42 {
    case 4294967295:
        fmt.Println(v7)
    default:
        if true {
            var v8 int = -2147483648
            fmt.Println(v2, v8, v7)
            _ = v8
        } else {
            if (65535 <= uint32(v1)) {
                for fi9 := 0; fi9 < 13; fi9++ {
                    fmt.Println(v3, v1)
                    _ = fi9
                }
                v3 = v3
                var v10 float32 = (1e10 + 1e10)
                _ = v10
            }
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 string = ((("hello" + ("\n\t" + "foo bar")) + "\n\t") + ("a" + ("a" + "a")))
                v13 := GenS1{F0: 1}
                v14 := GenS1{F0: 32767}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            v4 = append(v4, (("a" + ("foo bar" + "")) + ""))
        }
    }
    v5 = append(v5, "")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
