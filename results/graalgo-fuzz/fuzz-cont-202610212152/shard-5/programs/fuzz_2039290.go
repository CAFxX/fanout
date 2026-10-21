package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int64, p1 int64, p2 string) int64 {
    if (!((p1 / (((p0 & p1)) | 1)) < p0)) {
        p1 = (p1 - p0)
    }
    fmt.Println(p1, p0, p0)
    if (0 != 1) {
        switch ((p2 + ("foo bar" + "foo bar")) + (((("\n\t" + (p2 + (("a" + p2) + p2))) + p2) + "hello") + "")) {
        case "\n\t":
            for fi4 := 0; fi4 < 1; fi4++ {
                var v5 int32 = -2147483648
                var v6 uint = 151
                _ = fi4
                _ = v5
                _ = v6
            }
        }
        v7 := []int{}
        _ = v7
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return (p1 % (((p1 % (((p1 - (p1 << ((189) & 63)))) | 1))) | 1))
}

func main() {
    var v1 uint = 2
    v2 := []string{"hello", (("\n\t" + "\n\t") + ("hello" + ("hello" + "hello"))), "", ("hello" + ("a" + ("hello" + "foo bar")))}
    var v3 int8 = -128
    var v8 uint = (v1 >> (((^(v1))) & 63))
    var v9 int64 = genFn0(514, 0, (("\n\t" + ("a" + "")) + "\n\t"))
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v10 int16 = -32768
    for fi11 := 0; fi11 < 7; fi11++ {
        _ = len(v2)
        v12 := GenS1{F0: -9223372036854775808}
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: (v9 << ((42) & 63))}
    v2 = append(v2, "hello")
    var v14 string = "a"
    v9 += v9
    v15 := GenS1{F0: v9}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
